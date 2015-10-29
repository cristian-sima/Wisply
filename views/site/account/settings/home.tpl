<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/account">Account</a></li>
      <li class="active">Settings</li>
    </ul>
  </div>
  <div class="panel-body">
    <div>
      <h2>Delete account</h2>
      <div class="well">
        Please note that you cannot recover the account. <br/>
        <a class="btn btn-danger" id="deleteAccountButton">Delete my account</a>
      </div>
    </div>
  </div>
</div>
<script src="/static/js/account/settings/home.js"></script>
<script>
$(document).ready(function() {
  var module = wisply.getModule("account-settings"),
    page = new module.Page();
    page.init();
});
</script>
