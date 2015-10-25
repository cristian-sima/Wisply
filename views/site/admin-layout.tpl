<!DOCTYPE html>
<html lang="en">
{{ template "header" . }}
<body id="wisply-body">
  {{ template "javascript" . }}
  {{ template "menu" . }}
  <div class="container" >
    <div class="page-header">
      <div class="row">
        <div class="admin-sidebar col-sm-2 col-xs-2">
          {{ template "admin-menu" }}
        </div>
        <div class="col-sm-10 col-xs-10">
          {{.LayoutContent}}
        </div>
      </div>
    </div>
    {{ template "footer" .}}
  </div>
</body>
</html>
