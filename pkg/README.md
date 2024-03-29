# Go API client for client

Wodby Developer Documentation https://wodby.com/docs/dev

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 3.0.18
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./client"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.wodby.com/api/v3*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApplicationApi* | [**CreateApp**](docs/ApplicationApi.md#createapp) | **Post** /apps | 
*ApplicationApi* | [**DeleteApp**](docs/ApplicationApi.md#deleteapp) | **Delete** /apps/{id} | 
*ApplicationApi* | [**GetApp**](docs/ApplicationApi.md#getapp) | **Get** /apps/{id} | 
*ApplicationApi* | [**GetAppDrushAliases**](docs/ApplicationApi.md#getappdrushaliases) | **Get** /apps/{id}/drush-aliases | 
*ApplicationApi* | [**GetApps**](docs/ApplicationApi.md#getapps) | **Get** /apps | 
*BackupApi* | [**GetBackup**](docs/BackupApi.md#getbackup) | **Get** /backups/{id} | 
*BackupApi* | [**GetBackups**](docs/BackupApi.md#getbackups) | **Get** /backups | 
*DomainApi* | [**GetDomain**](docs/DomainApi.md#getdomain) | **Get** /domains/{id} | 
*DomainApi* | [**GetDomains**](docs/DomainApi.md#getdomains) | **Get** /domains | 
*GitRepositoryApi* | [**GetGitRepo**](docs/GitRepositoryApi.md#getgitrepo) | **Get** /git-repo/{id} | 
*GitRepositoryApi* | [**GetGitRepos**](docs/GitRepositoryApi.md#getgitrepos) | **Get** /git-repo | 
*InstanceApi* | [**CreateInstance**](docs/InstanceApi.md#createinstance) | **Post** /instances | 
*InstanceApi* | [**DeleteInstance**](docs/InstanceApi.md#deleteinstance) | **Delete** /instances/{id} | 
*InstanceApi* | [**DeployInstance**](docs/InstanceApi.md#deployinstance) | **Post** /instances/{id}/deploy | 
*InstanceApi* | [**DeployInstanceCodebase**](docs/InstanceApi.md#deployinstancecodebase) | **Post** /instances/{id}/deploy-codebase | 
*InstanceApi* | [**GetInstance**](docs/InstanceApi.md#getinstance) | **Get** /instances/{id} | 
*InstanceApi* | [**GetInstances**](docs/InstanceApi.md#getinstances) | **Get** /instances | 
*InstanceApi* | [**UpgradeInstance**](docs/InstanceApi.md#upgradeinstance) | **Post** /instances/{id}/upgrade | 
*InstanceApi* | [**UpgradeInstances**](docs/InstanceApi.md#upgradeinstances) | **Post** /instances/upgrade | 
*OrganizationApi* | [**GetOrg**](docs/OrganizationApi.md#getorg) | **Get** /orgs/{id} | 
*OrganizationApi* | [**GetOrgs**](docs/OrganizationApi.md#getorgs) | **Get** /orgs | 
*ServerApi* | [**GetServer**](docs/ServerApi.md#getserver) | **Get** /servers/{id} | 
*ServerApi* | [**GetServers**](docs/ServerApi.md#getservers) | **Get** /servers | 
*StackApi* | [**GetStack**](docs/StackApi.md#getstack) | **Get** /stacks/{id} | 
*StackApi* | [**GetStacks**](docs/StackApi.md#getstacks) | **Get** /stacks | 
*StackApi* | [**UpdateStackFromUpstream**](docs/StackApi.md#updatestackfromupstream) | **Post** /stacks/{id}/update | 
*StackApi* | [**UpdateStacksFromUpstream**](docs/StackApi.md#updatestacksfromupstream) | **Post** /stacks/update | 
*TaskApi* | [**GetTask**](docs/TaskApi.md#gettask) | **Get** /tasks/{id} | 
*TaskApi* | [**GetTasks**](docs/TaskApi.md#gettasks) | **Get** /tasks | 
*UserApi* | [**GetAuthenticatedUser**](docs/UserApi.md#getauthenticateduser) | **Get** /user | 


## Documentation For Models

 - [App](docs/App.md)
 - [Backup](docs/Backup.md)
 - [BackupFile](docs/BackupFile.md)
 - [BackupFiles](docs/BackupFiles.md)
 - [Domain](docs/Domain.md)
 - [GitRepo](docs/GitRepo.md)
 - [Instance](docs/Instance.md)
 - [InstanceType](docs/InstanceType.md)
 - [Org](docs/Org.md)
 - [RequestAppCreate](docs/RequestAppCreate.md)
 - [RequestAppCreateGit](docs/RequestAppCreateGit.md)
 - [RequestAppCreateServices](docs/RequestAppCreateServices.md)
 - [RequestInstanceCreate](docs/RequestInstanceCreate.md)
 - [RequestInstanceCreateGit](docs/RequestInstanceCreateGit.md)
 - [RequestInstanceDeploy](docs/RequestInstanceDeploy.md)
 - [RequestInstanceDeployCodebase](docs/RequestInstanceDeployCodebase.md)
 - [RequestInstancesUpgrade](docs/RequestInstancesUpgrade.md)
 - [RequestStacksUpdate](docs/RequestStacksUpdate.md)
 - [ResponseTask](docs/ResponseTask.md)
 - [ResponseTaskApp](docs/ResponseTaskApp.md)
 - [ResponseTaskInstance](docs/ResponseTaskInstance.md)
 - [Server](docs/Server.md)
 - [Stack](docs/Stack.md)
 - [StackService](docs/StackService.md)
 - [StackServiceImplementation](docs/StackServiceImplementation.md)
 - [Task](docs/Task.md)
 - [User](docs/User.md)


## Documentation For Authorization

## ApiKeyAuth
- **Type**: API key 

Example
```
	auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
		Key: "APIKEY",
		Prefix: "Bearer", // Omit if not necessary.
	})
    r, err := client.Service.Operation(auth, args)
```

## Author



