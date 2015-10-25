<!DOCTYPE html>
<html lang="en">
{{ template "header" .}}
<body id="wisply-body">
  {{ template "javascript" . }}
  {{ template "menu" . }}
  <div class="container" >
    {{.LayoutContent}}
    {{ template "footer" .}}
  </div>
</body>
</html>
