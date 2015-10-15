<!DOCTYPE html>
<html lang="en" hola_ext_inject="disabled">
{{ template "header" .}}
<body>
  {{ template "javascript" . }}
  {{ template "menu" . }}
  <div class="container" >
    {{.LayoutContent}}
    {{ template "footer" .}}
  </div>
</body>
</html>
