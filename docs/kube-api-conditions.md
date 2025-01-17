# Hive Integration - Conditions

Conditions provide a standard mechanism for higher-level status reporting from a controller. Read more about conditions [here](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties)

## AgentClusterInstall Conditions

AgentClusterInstall supported condition types are: `SpecSynced`, `RequirementsMet`, `Completed`, `Failed`, `Stopped` and `Validated`.

|Type|Status|Reason|Message|Description|
|----|----|-----|-------------------|-------------------|
|SpecSynced|True|SyncOK|The Spec has been successfully applied|If the Spec was successfully applied|
|SpecSynced|False|BackendError|The Spec could not be synced due to backend error: <err>|If the Spec was not applied due to 500 error|
|SpecSynced|False|InputError|The Spec could not be synced due to an input error: <err>|If the Spec was not applied due to 40X error|
||||||
|Validated|True|ValidationsPassing|The cluster's validations are passing|Otherwise than other conditions|
|Validated|False|ValidationsFailing|The cluster's validations are failing: "summary of failed validations"|If the cluster status is "insufficient" or "pending-for-input"|
|Validated|Unknown|ValidationsUnknown|The cluster's validations have not yet been calculated|If the validations have not yet been calculated|
||||||
|RequirementsMet|True|ClusterIsReady|The cluster is ready to begin the installation|if the cluster status is "ready"|
|RequirementsMet|False|ClusterNotReady|The cluster is not ready to begin the installation|If the cluster is before installation ("insufficient"/"pending-for-input")|
|RequirementsMet|True|ClusterAlreadyInstalling|The cluster requirements are met|If the cluster has begun installing ("preparing-for-installation", "installing", "finalizing", "installing-pending-user-action", "adding-hosts") |
|RequirementsMet|True|ClusterInstallationStopped|The cluster installation stopped|If the cluster has stopped installing ("installed", "error") |
|RequirementsMet|False|InsufficientAgents|The cluster currently requires `X` agents but only `Y` are ready|If the cluster is ready but we don't have the expected number of ready agents |
|RequirementsMet|False|UnapprovedAgents|The installation is pending on the approval of `X` agents|If the cluster is ready with the expected number of ready agents, but not all have been approved |
||||||
|Completed|True|InstallationCompleted|The installation has completed: "status_info"|If the cluster status is "installed"|
|Completed|False|InstallationFailed|The installation has failed: "status_info"|If the cluster status is "error"|
|Completed|False|InstallationNotStarted|The installation has not yet started|If the cluster is before installation ("insufficient"/"pending-for-input"/"ready")|
|Completed|False|InstallationInProgress|The installation is in progress: "status_info"|If the cluster is installing ("preparing-for-installation", "installing", "finalizing", "installing-pending-user-action")|
||||||
|Failed|True|InstallationFailed|The installation failed: "status_info"|if the cluster status is "error"|
|Failed|False|InstallationNotFailed|The installation has not failed|If the cluster status is not "error"|
||||||
|Stopped|True|InstallationFailed|The installation has stopped due to error|if the cluster status is "error"|
|Stopped|True|InstallationCancelled|The installation has stopped because it was cancelled|if the cluster status is "cancelled"|
|Stopped|True|InstallationCompleted|The installation has stopped because it completed successfully|if the cluster status is "installed"|
|Stopped|False|InstallationNotStopped|The installation is waiting to start or in progress|If the cluster status is not "error", "cancelled" or "installed|

Here an example of AgentClusterInstall conditions:

```sh
Status:
  Conditions:
    Last Probe Time:             2021-05-12T09:06:30Z
    Last Transition Time:        2021-05-12T09:06:30Z
    Message:                     The Spec has been successfully applied
    Reason:                      SyncOK
    Status:                      True
    Type:                        SpecSynced
    Last Probe Time:             2021-05-12T09:07:39Z
    Last Transition Time:        2021-05-12T09:07:39Z
    Message:                     The cluster requirements are met
    Reason:                      ClusterAlreadyInstalling
    Status:                      True
    Type:                        RequirementsMet
    Last Probe Time:             2021-05-12T09:07:39Z
    Last Transition Time:        2021-05-12T09:07:39Z
    Message:                     The cluster's validations are passing
    Reason:                      ValidationsPassing
    Status:                      True
    Type:                        Validated
    Last Probe Time:             2021-05-12T09:20:09Z
    Last Transition Time:        2021-05-12T09:20:09Z
    Message:                     The installation is in progress: Finalizing cluster installation
    Reason:                      InstallationInProgress
    Status:                      False
    Type:                        Completed
    Last Probe Time:             2021-05-12T09:06:30Z
    Last Transition Time:        2021-05-12T09:06:30Z
    Message:                     The installation has not failed
    Reason:                      InstallationNotFailed
    Status:                      False
    Type:                        Failed
    Last Probe Time:             2021-05-12T09:06:30Z
    Last Transition Time:        2021-05-12T09:06:30Z
    Message:                     The installation is waiting to start or in progress
    Reason:                      InstallationNotStopped
    Status:                      False
    Type:                        Stopped
```

## Agent Conditions

The Agent condition types supported are: `SpecSynced`, `Connected`, `ReadyForInstallation`, `Validated` and `Installed`

|Type|Status|Reason|Message|Description|
|----|----|-----|-------------------|-------------------|
|SpecSynced|True|SyncOK|The Spec has been successfully applied|If the Spec was successfully applied|
|SpecSynced|False|BackendError|The Spec could not be synced due to backend error: <err>|If the Spec was not applied due to 500 error|
|SpecSynced|False|InputError|The Spec could not be synced due to an input error: <err>|If the Spec was not applied due to 40X error|
||||||
|Validated|True|ValidationsPassing|The agent's validations are passing|Otherwise than other conditions|
|Validated|False|ValidationsFailing|The agent's validations are failing: "summary of failed validations"|If the host status is "insufficient" or "pending-for-input"|
|Validated|Unknown|ValidationsUnknown|The agent's validations have not yet been calculated|If the validations have not yet been calculated|
||||||
|ReadyForInstallation|True|AgentIsReady|The agent is ready to begin the installation|If the host is approved and in status "known"|
|ReadyForInstallation|False|AgentNotReady|The agent is not ready to begin the installation|If the host is before installation ("discovering"/"insufficient"/"disconnected"/"pending-input")|
|ReadyForInstallation|False|AgentAlreadyInstalling|The agent cannot begin the installation because it has already started|If the agent has begun installing ("preparing-successful","preparing-for-installation", "installing") |
|ReadyForInstallation|False|AgentIsNotApproved|The agent is not approved|If the host is not approved and in status "known"|
|ReadyForInstallation|False|AgentAlreadyInstalling|The agent cannot begin the installation because it has already started|If the agent has begun installing ("preparing-successful","preparing-for-installation", "installing") |
|ReadyForInstallation|False|AgentInstallationStopped|The agent installation stopped|If the agent has stopped installing ("installed", "error") |
||||||
|Installed|True|InstallationCompleted|The installation has completed: "status_info"|If the host status is "installed"|
|Installed|False|InstallationFailed|The installation has failed: "status_info"|If the host status is "error"|
|Installed|False|InstallationNotStarted|The installation has not yet started|If the cluster is before installation ("discovering"/"insufficient"/"disconnected"/"pending-input/known")|
|Installed|False|InstallationInProgress|The installation is in progress: "status_info"|If the host is installing ("preparing-for-installation", "preparing-successful", "installing")|
||||||
|Connected|True|AgentIsConnected|The agent has not contacted the installation service in some time, user action should be taken|If the host status is not "disconnected"|
|Connected|False|AgentIsDisconnected|The agent's connection to the installation service is unimpaired|If the host status is "error"|


Here an example of Agent conditions:

```sh
Status:
  Conditions:
    Last Transition Time:  2021-04-22T15:50:24Z
    Message:               The Spec has been successfully applied
    Reason:                SyncOK
    Status:                True
    Type:                  SpecSynced
    Last Transition Time:  2021-04-22T15:50:24Z
    Message:               The agent's connection to the installation service is unimpaired
    Reason:                AgentIsConnected
    Status:                True
    Type:                  Connected
    Last Transition Time:  2021-04-22T15:50:33Z
    Message:               The agent cannot begin the installation because it has already started
    Reason:                AgentAlreadyInstalling
    Status:                False
    Type:                  ReadyForInstallation
    Last Transition Time:  2021-04-22T15:50:26Z
    Message:               The agent's validations are passing
    Reason:                ValidationsPassing
    Status:                True
    Type:                  Validated
    Last Transition Time:  2021-04-22T15:50:24Z
    Message:               The installation is in progress: Host is preparing for installation
    Reason:                InstallationInProgress
    Status:                False
    Type:                  Installed
```

## InfraEnv Conditions

The InfraEnv condition type supported is: `ImageCreated`

|Type|Status|Reason|Message|Description|
|----|----|-----|-------------------|-------------------|
|ImageCreated|True|ImageCreated|Image has been created|If the ISO image was successfully created|
|ImageCreated|False|ImageCreationError|Failed to create image: "error message"|If the ISO image was not successfully created|

Here an example of InfraEnv conditions:

```sh
Status:
  Conditions:
    Last Transition Time:  2021-04-22T15:49:35Z
    Message:               Image has been created
    Reason:                ImageCreated
    Status:                True
    Type:                  ImageCreated
```
