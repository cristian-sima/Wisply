<div class="page-header" id="banner">
  <div class="row" >
    <div class="col-md-6 col-md-offset-3 col-centered" >
      <div class="panel panel-primary">
        <div class="panel-heading">
          <h1 class="panel-title">Register page</h1>
        </div>
        <div class="panel-body">
          <p>
            <div class=" text-center">
              We promise to respect your privacy.
            </div>
            <form action="/auth/register" method="POST" class="form-horizontal" id="register-form" >
              {{ .xsrf_input }}
              <div class="form-group text-left">
                <label for="register-name" class="col-lg-2 control-label">Name</label>
                <div class="col-lg-10">
                  <input type="text" value="{{.name}}" class="form-control" name="register-name" id="register-name" placeholder="Full name" required pattern="[A-Za-z0-9\s\.]{3,25}" title="The name has 3 up to 25 characters, can contain numbers, letters, spaces and dots.">
                </div>
              </div>
              <div class="form-group text-left">
                <label for="register-password" class="col-lg-2 control-label">Password</label>
                <div class="col-lg-10">
                  <input type="password" value="{{.passowrd}}" class="form-control" name="register-password" id="register-password" placeholder="Password" required pattern=".{6,25}" title="The password has 6 up to 25 characters!">
                </div>
              </div>
              <div class="form-group text-left" id="div-confirm-password" style="display:none">
                <label for="register-password-confirm" class="col-lg-2 control-label">Confirm Password</label>
                <div class="col-lg-10">
                  <input type="password" value="{{.passowrdConfirm}}" class="form-control" name="register-password-confirm" id="register-password-confirm" placeholder="Type the password again" required pattern=".{6,25}" title="The password has 6 up to 25 characters!">
                </div>
              </div>
              <div class="form-group text-left">
                <label for="register-email" class="col-lg-2 control-label">E-mail</label>
                <div class="col-lg-10">
                  <input type="email" value="{{.email}}" class="form-control" name="register-email" id="register-email" placeholder="E-mail address" required pattern=".{3,60}" title="You should provide a valid e-mail address.">
                </div>
              </div>
              {{ if .showCaptcha }}
              <div class="form-group text-left">
                <label for="register-captcha-value" class="col-lg-2 control-label">Verification</label>
                <div class="col-lg-10">
                  <div id="register-form-page-captcha">
                  </div>
                  <input type="text" class="form-control" placeholder="Type the numbers" name="register-form-page-captcha-value" id="register-captcha-value" required  />
                </div>
              </div>
              {{ end }}
              <div class="form-group" >
                <div class="text-center" id="register-submit-div">
                  <p class="text-muted">By creating this account, you agree to respect our <a target="_blank" href="/help#legal-aspects">legal aspects</a>.</p>
                  <input type="submit" class="btn btn-primary" value="Create account" />
                </div>
              </div>
            </form>
          </div>
        </div>
        <div class="form-group">
          <div class="panel panel-default">
            <div class="panel-body">
              Go back to the <a href="/auth/login"> Login</a> form.<br />
              <!--Do you need to recover your details? <a href="/auth/recover">Recover</a>-->
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <script src="/static/js/wisply/server.js"></script>

  {{ if .showCaptcha }}

  <script src="/static/js/wisply/captcha.js"></script>
  <script>
  $(document).ready(function() {
      var data = {
        hasCaptcha: true,
        ID: "{{ .captcha.GetID }}",
      };
      wisply.getModule("server").set(data);
  });
  </script>

  {{ else }}

  <script>
  $(document).ready(function() {
      var data = {
        hasCaptcha: false,
      };
      wisply.getModule("server").set(data);
  });
  </script>

  {{ end }}

  <script src="/static/js/public/auth/register.js"></script>
  <script>
  $(document).ready(function(){
      var module = wisply.getModule("register"),
        form = new module.Form();
        form.init();
  });
  </script>
