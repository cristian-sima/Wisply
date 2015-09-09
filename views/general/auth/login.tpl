<div class="page-header" id="banner">
  <div class="row" >
    <div class="col-md-6 col-md-offset-3 col-centered" >
      <div class="panel panel-primary">
        <div class="panel-heading">
         <h3 class="panel-title">Authentication</h3>
       </div>
        <div class="panel-body">
          <p>
            <div class=" text-center">
               We need you to connect
            </div>
            <form action="{{.actionURL}}" method="{{.actionType}}" class="form-horizontal" >
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
                    <input type="text" value="{{.passowrd}}" class="form-control" name="login-password" id="login-password" placeholder="Password" required pattern=".{3,25}" title="The password has 3 up to 25 characters!">
                  </div>
                </div>
                <div class="form-group">
                  <div class="text-center">
                    <input type="submit" class="btn btn-primary" href="#" role="button" value="Login" />
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
          Do you need an account?  <a href="/auth/register">Register</a> <br />
          Do you need to recover your details? <a href="/auth/recover">Recover</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
<script>
$(document).ready(function() {
    $("#login-username").focus();
});
</script>
