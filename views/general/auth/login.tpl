<div class="col-lg-8 col-md-7 col-sm-6" >
    <div class="panel panel-default">
        <div class="panel-body">
            <p>
              In order to visit this page, you must firstly connect.
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
                        <label for="password" class="col-lg-2 control-label">Password</label>
                        <div class="col-lg-10">
                            <input type="text" value="{{.passowrd}}" class="form-control" name="login-password" id="login-password" placeholder="Password" required pattern=".{3,25}" title="The password has 3 up to 25 characters!">
                        </div>
                    </div>
                    <div class="text-center">
                        <input type="submit" class="btn btn-default" href="#" role="button" value="Login" />
                    </div>
                </fieldset>
            </form>
            </p>
        </div>
    </div>
</div>

<div class = "container">
	<div class="wrapper">
		<form action="" method="post" name="Login_Form" class="form-signin">
		    <h3 class="form-signin-heading">We need your confirmation:</h3>
			  <hr class="colorgraph"><br>

			  <input type="text" class="form-control" name="Username" placeholder="Username" required="" autofocus="" />
			  <input type="password" class="form-control" name="Password" placeholder="Password" required=""/>

			  <button class="btn btn-lg btn-primary btn-block"  name="Submit" value="Login" type="Submit">Login</button>
		</form>
	</div>
</div>

<script>
$(document).ready(function() {
    $("#username").focus();
});
</script>
