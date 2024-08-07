# Application Groups

## Introduction

Application groups in Devtron streamline the deployment of microservices by enabling you to build and deploy multiple applications simultaneously. This feature is particularly beneficial when your microservices are interdependent, as a change in one service often triggers the need to redeploy others.

{% hint style="info" %}
Only one application group would exist for each [environment](../reference/glossary.md#environment). You cannot group applications belonging to different environments.
{% endhint %}

---

## Accessing Application Groups

1. From the left sidebar, go to **Application Groups**

    ![Figure 1: Application Group (Beta)](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/app-group-tab.jpg)

2. You will see a list of environments. Select the environment to view the application group.

    ![Figure 2: List of Environments](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/app-groups.jpg)

3. The application group would contain the applications meant for deployment in the chosen environment.

    ![Figure 3: Sample Application Group](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/app-group-overview-1.jpg)

As you can see, it has similar options as available under [Applications](./applications.md):
* Overview
* Build & Deploy
* Build history
* Deployment history
* Configurations

{% hint style="info" %}
Users need to have [View only permission](../user-guide/global-configurations/authorization/user-access.md#role-based-access-levels) or above (along with access to the environment and applications) to view all the applications within a group.
{% endhint %}

First, we will walk you through the [key features](#key-features) of Application Groups, followed by [additional features](#additional-features) that will help you perform bulk actions.

---

## Key Features

### Building Application Images

The **Build & Deploy** tab of your application group enables you to trigger the [CI builds](../reference/glossary.md#image) of one or more applications in bulk.

1. Select the applications using the checkboxes and click the **Build Image** button present at the bottom.

    ![Figure 4: Build Option](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/select-app.jpg)

2. The `Build image` screen opens. Select the application and the [commit](../reference/glossary.md#commit-hash) for which you want to trigger the CI build.

    ![Figure 5: Selecting Commit](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/select-commit-1.jpg)

{% hint style="info" %}
### Tip
Adding [image labels](./deploying-application/image-labels-and-comments.md) can help you quickly locate the container image from the list of images shown in Application Groups.
{% endhint %}

3. Similar to application, you can also [pass build parameters](./deploying-application/triggering-ci.md#passing-build-parameters) in application groups before triggering the build.

{% hint style="info" %}
Passing build parameters feature is only available in <img src="https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/elements/EnterpriseTag.svg">
{% endhint %}

* Go to the **Parameters** tab.

    ![Figure 6: Parameters Tab](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/ag-parameter-tab.jpg)

* Click **+ Add parameter**.

    ![Figure 7: Adding a Parameter](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/ag-add-parameter.jpg)

* Enter your key-value pair as shown below. 

    ![Figure 8: Entering Key-Value Pair](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/ag-key-value.jpg)

* You may follow the above steps for other applications too, and then click **Start Build**.

    ![Figure 9: Choosing Commit for Other Application](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/ag-next-app.jpg)

    ![Figure 10: Passing Build Parameters and Triggering Build](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/ag-start-build.jpg)

4. The builds will initiate, following which, you can close the `Build image` screen.

    ![Figure 11: Triggered Deployment](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/build-image.jpg)

{% hint style="info" %}
Users need to have [Build and deploy permission](../user-guide/global-configurations/authorization/user-access.md#role-based-access-levels) or above (along with access to the environment and applications) to trigger the build
{% endhint %}


### Changing Configurations

The **Configurations** tab of your application group allows you to configure the following:

* [Deployment template](../reference/glossary.md#base-deployment-template)
* [ConfigMaps](../reference/glossary.md#configmaps)
* [Secrets](../reference/glossary.md#secrets)

As shown below, you can handle the configurations of more than one application from a single screen.

![Figure 12: Configurations of each App](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/configurations.jpg)

{% hint style="info" %}
Users need to have [Admin role](../user-guide/global-configurations/authorization/user-access.md#role-based-access-levels) or above (along with access to the environment and applications) to change their configuration. Please note, you might not be able to change the values of locked keys in deployment template. Refer [Lock Deployment Configuration](./global-configurations/lock-deployment-config.md) to know more.
{% endhint %}


### Deploying Applications

The **Build & Deploy** tab of your application group helps you deploy one or more applications in bulk.

1. Select the applications using the checkboxes and click the **Deploy** button present at the bottom.

    ![Figure 13: Deploy Option](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/select-app-deploy.jpg)

2. Select the desired container image that you want to deploy for respective application.

    ![Figure 14: Selecting Image](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/select-image-1.jpg)

    Repeat the step for other applications too, and then click **Deploy**.

    ![Figure 15: Deploying Apps](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/select-image-2.jpg)

3. The deployment will be initiated, following which, you can close the screen as shown below.

    ![Figure 16: Triggered Deployment](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/deploy-app.jpg)

Once the deployment is successful, the pipelines will show `Succeeded`.

![Figure 17: Successful Deployment](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/successful.jpg)

{% hint style="info" %}
Users need to have [Build and deploy permission](../user-guide/global-configurations/authorization/user-access.md#role-based-access-levels) or above (along with access to the environment and applications) to initiate the deployment
{% endhint %}

---

## Additional Features

### Hibernating and Unhibernating Apps

{% hint style="warning" %}
### Who Can Perform This Action?
Users need to have [Build & deploy permission](./global-configurations/authorization/user-access.md#role-based-access-levels) or above (along with access to the environment and application) to hibernate or unhibernate applications.
{% endhint %}

Since every application comes with an option to hibernate, the same is true for application groups. Using application group, you can hibernate one or more applications belonging to the same environment if you do not want them to consume resources (replica count will be set to 0). 

In other words, you can hibernate running applications or unhibernate hibernated applications as per your requirement.

#### Hibernation Process

1. In the `Overview` page of your application group, use the checkboxes to choose the applications you wish to hibernate, and click the **Hibernate** button.

    ![Figure 18a: Selecting Apps to Hibernate](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/hibernate-apps-v1.jpg)

2. Confirm the hibernation.

    ![Figure 18b: Confirming Hibernation](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/confirm-hibernation-v1.jpg)

3. Hibernation will initiate as shown below. You may close the window. 

    ![Figure 18c: Initiation Status of Hibernation](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/initiated-hibernation.jpg)

Your applications pods would be scaled down and would stop incurring costs.

#### Unhibernation Process

1. In the same `Overview` page, you can use the checkboxes to choose the hibernated applications you wish to unhibernate, and click the **Unhibernate** button.

    ![Figure 19a: Selecting Hibernated Apps to Unhibernate](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/unhibernate-apps-v1.jpg)

2. Confirm the unhibernation.

    ![Figure 19b: Confirming Unhibernation](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/confirm-unhibernation-v1.jpg)

3. Unhibernation will initiate as shown below. You may close the window. 

    ![Figure 19c: Initiation Status of Unhibernation](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/initiated-unhibernation.jpg)

Your applications would be up and running in some time.

### Restart Workloads

{% hint style="warning" %}
### Who Can Perform This Action?
Users need to have [Build & deploy permission](./global-configurations/authorization/user-access.md#role-based-access-levels) or above (along with access to the environment and application) to restart workloads in bulk.
{% endhint %}

Restarting workloads might be necessary if you want your new code or configuration to come into effect, or you are experiencing issues like crashing of pods.  

Using application group, you can select the workloads (i.e., Pod, Deployment, ReplicaSet, etc.) of specific applications and restart them. 

1. Use the checkboxes to choose the applications whose workloads you wish to restart, and click the **Restart Workload** button.

    ![Figure 20a: Selecting Apps to Restart](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/restart-workloads-v1.jpg)

2. Next to the application, click the workload dropdown to view all the individual workloads of an application. Choose only the ones you wish to restart.

    ![Figure 20b: Choosing Workloads](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/choose-workload.jpg)

    Moreover, you can easily select, deselect, or choose multiple workloads as shown below.

    ![Figure 20c: Selecting and Unselecting Workloads](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/bulk-restart.gif)

3. Click **Restart Workloads**.

    ![Figure 20d: Restarting Workloads](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/select-workloads.jpg)

Restarting workloads might take time depending on the number of applications.

### Filtering Applications

Assume you have multiple applications (maybe 10, 50, 100, or more) showing up in an application group. If you want to limit your operations (build/deploy/other) to a specific set of applications, the filter feature will help you narrow down the list. Thus, you will see only those applications you select from the filter (be it on the `Overview` page, `Build & Deploy` page, and so on.)

1. Click the filter next to the application group as shown below.

    ![Figure 21: Filter Option](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/app-filter-1.jpg)

2. The filter will show all the applications present in the group. Click to select the relevant ones.

    ![Figure 22: All Apps](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/app-filter-2.jpg)

3. The filter narrows down the list of applications as shown below.

    ![Figure 23: Filtered Apps](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/app-filter-3.jpg)

4. (Optional) If required, you can save the filter for future use by clicking **Save selection as filter**.

    ![Figure 24: Saving a Filter](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/save-filter.jpg)

5. Add a name and description to the filter to help you know its purpose, and click **Save**.

    ![Figure 25: Naming a Filter](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/save-filter-2.jpg)

Now when you access the application group, your saved filter will be visible on top.

![Figure 26: Saved Filter](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/save-filter-3.jpg)

{% hint style="info" %}

### Permissions

#### 1. Creating a filter

Users can create a filter if they have Admin/Manager access on all selected applications.

* **Case 1**: User has Admin/Manager access on all selected applications

    User will be able to create a filter with all selected applications.

* **Case 2**: User does not have Admin/Manager access on all selected applications

    User will not be able to create a filter.

* **Case 3**: User selected 4 applications but has Admin/Manager access for only 2 of them

    User should be able to create filter with these 2 applications.

#### 2. Editing a saved filter

Users can edit a saved filter if they have Admin/Manager access on all applications in the saved filter.

#### 3. Deleting a saved filter

Users can delete a saved filter if they have Admin/Manager access on all applications in the saved filter.

{% endhint %}


### Changing Branch

{% hint style="warning" %}
### Who Can Perform This Action?
Users need to have [Admin role](../user-guide/global-configurations/authorization/user-access.md#role-based-access-levels) or above (along with access to the environment and applications) to update their branch.
{% endhint %}

Assume you have a few applications whose [build pipelines](../reference/glossary.md#build-pipeline) fetch from the `main` branch of your code repository. However, you decided to maintain a `master` branch, and you want all the upcoming CI builds to consider the `master` branch as the source. Devtron provides you the option to change the branch at both levels—individual application as well as application group.

1. In the **Build & Deploy** tab of your application group, select the intended applications and click the **Change Branch** button present at the bottom.

    ![Figure 27: Changing Branch](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/change-branch.jpg)

2. Enter the new branch name. If your build pipeline has `Branch Regex` as the Source Type, you must ensure your new branch name matches the regex (regular expression) provided in that build pipeline. Once done, click **Update Branch**.

    ![Figure 28: Updating Branch Name](https://devtron-public-asset.s3.us-east-2.amazonaws.com/images/application-groups/update-branch.jpg)
