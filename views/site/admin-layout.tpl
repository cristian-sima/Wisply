<!DOCTYPE html>
<html lang="en" hola_ext_inject="disabled">
{{ template "header" . }}
<body>
  {{ template "javascript" . }}
  {{ template "menu" . }}
  <div class="container" >
    <div class="page-header" id="banner">
      <div class="row row-offcanvas row-offcanvas-left">
        <div class="col-xs-2 col-sm-2 sidebar-offcanvas" id="sidebar" role="navigation">
          {{ template "admin-menu" }}
        </div>
       <div class="col-xs-12 col-sm-10">
          {{.LayoutContent}}
        </div>
      </div>
    </div>
    {{ template "footer" .}}
  </div>
</body>
</html>
