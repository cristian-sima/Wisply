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
            <form action="/auth/register" method="POST" class="form-horizontal" id="register-form" >
              {{ .xsrf_input }}
              <fieldset>
                <div class="form-group">
                  <label for="register-username" class="col-lg-2 control-label">Username</label>
                  <div class="col-lg-10">
                    <input type="text" value="{{.username}}" class="form-control" name="register-username" id="register-username" placeholder="Username" required pattern=".{3,25}" title="The username has 3 up to 25 characters!">
                  </div>
                </div>
                <div class="form-group">
                  <label for="register-password" class="col-lg-2 control-label">Password</label>
                  <div class="col-lg-10">
                    <input type="password" value="{{.passowrd}}" class="form-control" name="register-password" id="register-password" placeholder="Password" required pattern=".{6,25}" title="The password has 6 up to 25 characters!">
                  </div>
                </div>
                <div class="form-group" id="div-confirm-password" style="display:none">
                  <label for="register-password-confirm" class="col-lg-2 control-label">Confirm Password</label>
                  <div class="col-lg-10">
                    <input type="password" value="{{.passowrdConfirm}}" class="form-control" name="register-password-confirm" id="register-password-confirm" placeholder="Type the password again" required pattern=".{6,25}" title="The password has 6 up to 25 characters!">
                  </div>
                </div>
                <div class="form-group">
                  <label for="register-email" class="col-lg-2 control-label">E-mail</label>
                  <div class="col-lg-10">
                    <input type="email" value="{{.email}}" class="form-control" name="register-email" id="register-email" placeholder="E-mail address" required pattern=".{3,25}" title="You should provide a valid e-mail address.">
                  </div>
                </div>
                <div class="form-group" >
                  <div class="text-center" id="register-submit-div">
                    <input type="submit" class="btn btn-primary" href="#" role="button" value="Create account" />
                  </div>
                </div>
              </fieldset>
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
<script>
$(document).ready(function() {
  $("#register-username").focus();
  addListeners();
  checkConfirmPassword();
});
function checkConfirmPassword () {
  if($("#register-password-confirm").val() !== "") {
    $("#div-confirm-password").show();
  }
}
function addListeners () {
  passwordCompletedListener();
  registerForm();
}
function passwordCompletedListener() {
  $("#register-password").focus(function() {
    $("#div-confirm-password").show();
  });
}
function registerForm() {
  $("#register-form").on("submit", registerFormSubmited)
}
function registerFormSubmited(e) {
  e.preventDefault();
  var password = $('#register-password').val(),
  confirmationPassword = $("#register-password-confirm").val();
  if (password === confirmationPassword) {
    showLoading();
    this.submit();
  } else {
    alertUser();
  }
}
function showLoading() {
  setElementLoading($('#register-submit-div'), "medium")
}

function alertUser() {
  bootbox.dialog({
    title: "I got a problem",
    message: "The confirmation password is equal to the password. Correct this.",
    onEscape: true,
    buttons: {
      cancel: {
        label: "Ok",
        className: "btn-default",
        callback: function() {
          this.modal('hide');
        }
      }}
    });
  }

  </script>
