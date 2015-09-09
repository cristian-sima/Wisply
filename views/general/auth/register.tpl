<div class="page-header" id="banner">
  <div class="row" >
    <div class="col-md-6 col-md-offset-3 col-centered" >
      <div class="panel panel-primary">
        <div class="panel-heading">
         <h3 class="panel-title">Register</h3>
       </div>
        <div class="panel-body">
          <p>
            <div class=" text-center">
               We like privacy. We promise.
            </div>
            <form action="" method="POST" class="form-horizontal" >
              {{ .xsrf_input }}
              <fieldset>
                <legend>{{.legend}}</legend>
                <div class="form-group">
                  <label for="login-username" class="col-lg-2 control-label">Username</label>
                  <div class="col-lg-10">
                    <input type="text" value="{{.username}}" class="form-control" name="login-username" id="login-username" placeholder="Username" required pattern=".{3,25}" title="The username has 3 up to 25 characters!">
                  </div>
                </div>
                <div class="form-group">
                  <label for="login-password" class="col-lg-2 control-label">Password</label>
                  <div class="col-lg-10">
                    <input type="password" value="{{.passowrd}}" class="form-control" name="login-password" id="login-password" placeholder="Password" required pattern=".{6,25}" title="The password has 6 up to 25 characters!">
                  </div>
                </div>
                <div class="form-group" id="div-confirm-password" style="display:none">
                  <label for="login-password-confirm" class="col-lg-2 control-label">Confirm Password</label>
                  <div class="col-lg-10">
                    <input type="password" value="{{.passowrdConfirm}}" class="form-control" name="login-password-confirm" id="login-password-confirm" placeholder="Type the password again" required pattern=".{6,25}" title="The password has 6 up to 25 characters!">
                  </div>
                </div>
                <div class="form-group">
                  <label for="login-email" class="col-lg-2 control-label">E-mail</label>
                  <div class="col-lg-10">
                    <input type="email" value="{{.email}}" class="form-control" name="login-email" id="login-email" placeholder="E-mail address" required pattern=".{3,25}" title="You should provide a valid e-mail address.">
                  </div>
                </div>
                <div class="form-group">
                  <div class="text-center">
                    <input type="submit" class="btn btn-primary" href="#" role="button" value="Create account" />
                  </div>
                </div>
              </fieldset>
            </form>
          </p>
        </div>
      </div>
      <div class="form-group">
        <div class="panel panel-default">
          <div class="panel-body">
          Go back to the <a href="/auth/login"> Login</a> form.<br />
          Do you need to recover your details? <a href="/auth/recover">Recover</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
<script>
function checkConfirmPassword () {
    if($("#login-password-confirm").val() !== "") {
      $("#div-confirm-password").show();
    }
}
function addListener () {
  $("#login-password").focus(function() {
    $("#div-confirm-password").show();
  })
}
$(document).ready(function() {
    $("#login-username").focus();
    addListener();
    checkConfirmPassword();
});
</script>
