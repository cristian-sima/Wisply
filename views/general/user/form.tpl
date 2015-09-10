<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/users">Users</a></li>
      <li class="active">Modify</li>
    </ul></div>
    <div class="panel-body">
      <p>
        <form action="" method="POST" class="form-horizontal" id="modify">
          {{ .xsrf_input }}
          <fieldset>
            <legend>Modify <span class="label label-default">{{ .userUsername }}</span></legend>
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
                <button type="submit" class="btn btn-primary">Modify</button><a href="/admin/users"> <button type="button" class="btn btn-default">Cancel</button></a>
              </div>
            </div>
          </fieldset>
        </form>
      </p>
    </div>
  </div>
  <script>
  $(document).ready(function() {
    $("#user-username").focus();
  });
  </script>
