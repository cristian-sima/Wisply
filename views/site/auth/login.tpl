<div class="page-header" id="banner">
  <div class="row" >
    <div class="col-md-6 col-md-offset-3 col-centered" >
      <div class="panel panel-primary">
        <div class="panel-heading">
          <h3 class="panel-title">Authentication</h3>
        </div>
        <div class="panel-body">
          <div class=" text-center">
            <br />
          </div>
          <form action="/auth/login" method="POST" class="form-horizontal" id="login-form">
            {{ .xsrf_input }}
            <input type="hidden" value="{{.sendMe}}" name="login-send-me" />
            <fieldset>
              <div class="form-group text-left">
                <label for="login-email" class="col-lg-2 control-label">E-mail</label>
                <div class="col-lg-10">
                  <input type="email" class="form-control" name="login-email" id="login-email" placeholder="E-mail"  />
                </div>
              </div>
              <div class="form-group text-left">
                <label for="login-password" class="col-lg-2 control-label">Password</label>
                <div class="col-lg-10">
                  <input type="password" class="form-control" name="login-password" id="login-password" placeholder="Password" required pattern=".{3,25}" title="The password has 3 up to 25 characters!">
                  <div class="checkbox">
                    <label for="login-remember-me">
                      <input type="checkbox" name="login-remember-me" id="login-remember-me"> Remember me
                    </label>
                  </div>
                </div>
              </div>
              <div class="form-group">
                <div class="text-center" id="login-submit-div">
                  <input type="submit" class="btn btn-primary" id="login-submit-button" value="Login" />
                </div>
              </div>
            </fieldset>
          </form>
        </div>
      </div>
      <div class="form-group">
        <div class="panel panel-default">
          <div class="panel-body">
            Are you new? <a href="/auth/register">Register</a> in a few seconds <br />
            <!--Do you need to recover your details? <a href="/auth/recover">Recover</a>-->
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
<script src="/static/js/static/auth/login.js"></script>
