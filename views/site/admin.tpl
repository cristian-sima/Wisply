<!DOCTYPE html>
<html lang="en" hola_ext_inject="disabled">
    {{ template "header" }}
    <body>
        {{ template "javascript" . }}
        {{ template "menu" . }}
        <div class="container" >
            <div class="page-header" id="banner">
                <div class="row" >
                      <aside class="col-lg-2">
                        {{ template "admin-menu" }}
                      </aside>
                      <div class="col-lg-10">
                        {{.LayoutContent}}
                      </div>
                </div>
            </div>
            {{ template "footer"}}
        </div>
    </body>
</html>
