
<!-- .slide: class="title" -->
# Multi-Tenancy and RBAC

## Securing Access

<img src="/modules/orgs_and_acls/security.png" style="bottom: 00px" class="title-inset">

---

<!-- .slide: class="content small-font" -->
## Supporting Multiple Orgs

Velociraptor supports multiple orgs in a fully multi tenancy configuration.

* Clients are divided into `Orgs`
* Each `Org` is completely separated:
   * Files are stored in a different location
   * Users have different ACLs and rights in different orgs
   * Custom artifacts can be maintained in different Orgs
* Orgs can be created and destroyed easily at runtime
* All clients share the same infrastructure
* Different orgs' clients can not connect to other orgs

---

<!-- .slide: class="content small-font" -->
## Switching to different orgs

By default the `Velociraptor gui` command creates two orgs. Switch to
the second org sing the GUI's user preferences tab.

![](/modules/orgs_and_acls/selecting_orgs.png)


---

<!-- .slide: class="content small-font" -->
## Creating a new org

You can use the Server.Orgs.NewOrg artifact to create a new org

![](/modules/orgs_and_acls/new_org.png)

---

<!-- .slide: class="content small-font" -->

## Users and orgs

* A Velociraptor user is any entity that has permissions on the org
* Users may be able to log into the GUI or via the API (service accounts)
* Users need certain permissions to perform an action
* `Roles` are bundles of permissions - just a convenience! Extra
  permissions can also be given.

The default roles:
<div class="container small-font">
<div class="col">

* `org_admin`
* `administrator`
* `reader`
* `api`

</div>
<div class="col">

* `analyst`
* `investigator`
* `artifact_writer`

</div>
</div>

---

<!-- .slide: class="content small-font" -->

## Adding a new user

If using basic authentication you can change the user's password here
as well.

![](/modules/orgs_and_acls/add_new_user.png)


---

<!-- .slide: class="content small-font" -->

## Assign user to org

By default the initial role assigned is `reader`


![](/modules/orgs_and_acls/assign_user_to_org.png)


---

<!-- .slide: class="content small-font" -->

## Adjust User permissions

User roles and permissions are only effective within an org. The same
user can have different roles in different orgs.

To delete a user, simply remove all their roles from an org.

![](/modules/orgs_and_acls/adjust_user_permissions.png)

---

<!-- .slide: class="content small-font" -->

## Preparing a deployment for the new org

Clients are provisioned for their respective orgs.  You can prepare an
MSI for deployment using the `Server.Utils.CreateMSI` artifact.

![](/modules/orgs_and_acls/preparing_msi.png)


---

<!-- .slide: class="content small-font" -->

## Fetching the prepared MSI for deployment

The prepared MSI contains the relevant embedded config and is ready
for installation.

![](/modules/orgs_and_acls/download_msi.png)

---

<!-- .slide: class="content small-font" -->

## Auditing user action

Velociraptor is a very powerful platform and requires strong auditing.

![](/modules/orgs_and_acls/viewing_audit_log.png)

---

<!-- .slide: class="content small-font" -->

## Inspecting the audit timeline

It is also possible to forward audit events off system (remote syslog
or Open Search server)

![](/modules/orgs_and_acls/inspecting_audit_log.png)
