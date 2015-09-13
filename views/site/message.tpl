<!DOCTYPE html>
<html lang="en" hola_ext_inject="disabled">
{{ template "header" }}
<body>
  {{ template "javascript" }}
  {{ template "menu" }}
  <div class="container" >
    <div class="page-header" id="banner">
      <div class="row" >
        <div class="col-lg-12 col-md-12 col-sm-12" >
          {{.LayoutContent}}
          <br>
        </div>
      </div>
    </div>
    {{ template "footer"}}
  </div>
</body>
</html>
