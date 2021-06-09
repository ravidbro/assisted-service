package host

import (
	"github.com/filanov/stateswitch"
	"github.com/openshift/assisted-service/models"
)

func NewPoolClusterHostStateMachine(sm stateswitch.StateMachine, th *transitionHandler) stateswitch.StateMachine {
	// sm := stateswitch.NewStateMachine()

	sm.AddTransition(stateswitch.TransitionRule{
		TransitionType: TransitionTypeRegisterHost,
		SourceStates: []stateswitch.State{
			"",
			stateswitch.State(models.HostStatusWaitingToBeRegistered),
		},
		Condition:        th.IsPoolClusterHost,
		DestinationState: stateswitch.State(models.HostStatusDiscoveringPoolCluster),
		PostTransition:   th.PostRegisterHost,
	})

	// Register host
	sm.AddTransition(stateswitch.TransitionRule{
		TransitionType: TransitionTypeRegisterHost,
		SourceStates: []stateswitch.State{
			stateswitch.State(models.HostStatusDiscoveringPoolCluster),
			stateswitch.State(models.HostStatusDisconnectedPoolCluster),
			stateswitch.State(models.HostStatusInsufficientPoolCluster),
			stateswitch.State(models.HostStatusKnownPoolCluster),
		},
		DestinationState: stateswitch.State(models.HostStatusDiscoveringPoolCluster),
		PostTransition:   th.PostRegisterHost,
	})

	// Disabled host can register if it was booted, no change in the state.
	sm.AddTransition(stateswitch.TransitionRule{
		TransitionType:   TransitionTypeRegisterHost,
		SourceStates:     []stateswitch.State{stateswitch.State(models.HostStatusDisabledPoolCluster)},
		DestinationState: stateswitch.State(models.HostStatusDisabledPoolCluster),
	})

	// Disable host
	sm.AddTransition(stateswitch.TransitionRule{
		TransitionType: TransitionTypeDisableHost,
		SourceStates: []stateswitch.State{
			stateswitch.State(models.HostStatusDisconnectedPoolCluster),
			stateswitch.State(models.HostStatusDiscoveringPoolCluster),
			stateswitch.State(models.HostStatusInsufficientPoolCluster),
			stateswitch.State(models.HostStatusKnownPoolCluster),
		},
		DestinationState: stateswitch.State(models.HostStatusDisabledPoolCluster),
		PostTransition:   th.PostDisableHost,
	})

	// Enable host
	sm.AddTransition(stateswitch.TransitionRule{
		TransitionType: TransitionTypeEnableHost,
		SourceStates: []stateswitch.State{
			stateswitch.State(models.HostStatusDisabledPoolCluster),
		},
		DestinationState: stateswitch.State(models.HostStatusDiscoveringPoolCluster),
		PostTransition:   th.PostEnableHost,
	})

	// Bind host
	sm.AddTransition(stateswitch.TransitionRule{
		TransitionType: TransitionTypeBindHost,
		SourceStates: []stateswitch.State{
			stateswitch.State(models.HostStatusKnownPoolCluster),
		},
		DestinationState: stateswitch.State(models.HostStatusWaitingToBeRegistered),
		PostTransition:   th.PostBindHost,
	})

	// Refresh host

	sm.AddTransition(stateswitch.TransitionRule{
		TransitionType: TransitionTypeRefresh,
		SourceStates: []stateswitch.State{
			stateswitch.State(models.HostStatusDiscoveringPoolCluster),
			stateswitch.State(models.HostStatusInsufficientPoolCluster),
			stateswitch.State(models.HostStatusKnownPoolCluster),
			stateswitch.State(models.HostStatusDisconnectedPoolCluster),
		},
		Condition:        stateswitch.Not(If(IsConnected)),
		DestinationState: stateswitch.State(models.HostStatusDisconnectedPoolCluster),
		PostTransition:   th.PostRefreshHost(statusInfoDisconnected),
	})

	sm.AddTransition(stateswitch.TransitionRule{
		TransitionType: TransitionTypeRefresh,
		SourceStates: []stateswitch.State{
			stateswitch.State(models.HostStatusDisconnectedPoolCluster),
			stateswitch.State(models.HostStatusDiscoveringPoolCluster),
		},
		Condition:        stateswitch.And(If(IsConnected), stateswitch.Not(If(HasInventory))),
		DestinationState: stateswitch.State(models.HostStatusDiscoveringPoolCluster),
		PostTransition:   th.PostRefreshHost(statusInfoDiscovering),
	})

	var hasMinRequiredHardware = stateswitch.And(If(HasMinValidDisks), If(HasMinCPUCores), If(HasMinMemory), If(IsPlatformValid))

	// In order for this transition to be fired at least one of the validations in minRequiredHardwareValidations must fail.
	// This transition handles the case that a host does not pass minimum hardware requirements for any of the roles
	sm.AddTransition(stateswitch.TransitionRule{
		TransitionType: TransitionTypeRefresh,
		SourceStates: []stateswitch.State{
			stateswitch.State(models.HostStatusDisconnectedPoolCluster),
			stateswitch.State(models.HostStatusDiscoveringPoolCluster),
			stateswitch.State(models.HostStatusInsufficientPoolCluster),
			stateswitch.State(models.HostStatusKnownPoolCluster),
		},
		Condition: stateswitch.And(If(IsConnected), If(HasInventory),
			stateswitch.Not(hasMinRequiredHardware)),
		DestinationState: stateswitch.State(models.HostStatusInsufficientPoolCluster),
		PostTransition:   th.PostRefreshHost(statusInfoInsufficientHardware),
	})

	// Noop transitions
	for _, state := range []stateswitch.State{
		stateswitch.State(models.HostStatusDisabledPoolCluster),
		stateswitch.State(models.HostStatusWaitingToBeRegistered),
	} {
		sm.AddTransition(stateswitch.TransitionRule{
			TransitionType:   TransitionTypeRefresh,
			SourceStates:     []stateswitch.State{state},
			DestinationState: state,
		})
	}

	sm.AddTransition(stateswitch.TransitionRule{
		TransitionType: TransitionTypeRefresh,
		SourceStates: []stateswitch.State{
			stateswitch.State(models.HostStatusDisconnectedPoolCluster),
			stateswitch.State(models.HostStatusDiscoveringPoolCluster),
			stateswitch.State(models.HostStatusInsufficientPoolCluster),
			stateswitch.State(models.HostStatusKnownPoolCluster),
		},
		Condition: stateswitch.And(If(IsConnected), If(HasInventory),
			hasMinRequiredHardware),
		DestinationState: stateswitch.State(models.HostStatusKnownPoolCluster),
		PostTransition:   th.PostRefreshHost(statusInfoHostReadyToBeMoved),
	})

	return sm
}
