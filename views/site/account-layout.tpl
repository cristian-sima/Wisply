<!DOCTYPE html>
<html lang="en">
{{ template "header" . }}
<body>
  {{ template "javascript" . }}
  {{ template "menu" . }}
  <div class="container" >
    <div class="page-header" id="banner">
      {{ template "account-menu" }}
      <div class="row row-offcanvas row-offcanvas-left">
        <div class="col-xs-12 col-sm-12">
          {{.LayoutContent}}
        </div>
      </div>
    </div>
    {{ template "footer" .}}
  </div>
</body>
</html>
