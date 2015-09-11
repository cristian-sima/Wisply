<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/users">Users</a></li>
      <li class="active">Modify</li>
    </ul></div>
    <div class="panel-body">
        <form method="POST" class="form-horizontal" id="modify">
          {{ .xsrf_input }}
          <fieldset>
            <div class="form-group">
              <label for="modify-administrator" class="col-lg-2 control-label">Type</label>
              <div class="col-lg-10">
                <select name="modify-administrator" id="modify-administrator" class="form-control">
                  <option  value="true" {{ if .isAdministrator }} selected {{ end }} >Administrator</option>
                  <option  value="false" {{ if .isUser }} selected {{ end}}>User</option>
                </select>
              </div>
            </div>
            <div class="form-group">
              <div class="col-lg-10 col-lg-offset-2">
                <input type="submit" class="btn btn-primary" value="Modify"/> <a href="/admin/users" class="btn btn-default">Cancel</a>
              </div>
            </div>
          </fieldset>
        </form>
    </div>
  </div>
  <script>
  $(document).ready(function() {
    $("#user-username").focus();
  });
  </script>
