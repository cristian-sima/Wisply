<!DOCTYPE html>
<html lang="en" hola_ext_inject="disabled">
    {{ template "header" }}
    <body>
        {{ template "menu" }}
        <div class="container">
            {{.LayoutContent}}            
            {{ template "footer"}}
        </div>
        {{ template "javascript-bottom" }}
    </body>
</html>
