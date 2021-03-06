<div class="page-header" id="banner">
  <div class="row" >
    <div class="col-md-6 col-md-offset-3 col-centered" >
      <div class="panel panel-primary">
        <div class="panel-heading">
          <h1 class="panel-title">Authentication page</h1>
        </div>
        <div class="panel-body">
          <div class=" text-center">
            <br />
          </div>
          <form action="/auth/login" method="POST" class="form-horizontal" id="login-form">
            {{ .xsrf_input }}
            <input type="hidden" value="{{.sendMe}}" name="login-send-me" />
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
            {{ if .showCaptcha }}
            <div class="form-group text-left">
              <label for="login-captcha-value" class="col-lg-2 control-label">Verification</label>
              <div class="col-lg-10">
                <div id="login-form-page-captcha">
                </div>
                <input type="text" class="form-control" placeholder="Type the numbers" name="login-form-page-captcha-value" id="login-captcha-value" required  />
              </div>
            </div>
            {{ end }}
            <div class="form-group">
              <div class="text-center" id="login-submit-div">
                <input type="submit" class="btn btn-primary" id="login-submit-button" value="Login" />
              </div>
            </div>
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

<script src="/static/js/public/auth/login.js"></script>
<script>
$(document).ready(function(){
    var module = wisply.getModule("login"),
      form = new module.Form();
      form.init();
});
</script>
