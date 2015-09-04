{{ define "admin-menu" }}
<div class="col-lg-4 col-md-5 col-sm-6">
    <ul class="nav nav-pills nav-stacked">
        <li class="active"><a href="/admin">Home</a></li>
        <li class="dropdown">
            <a class="dropdown-toggle" data-toggle="dropdown" href="/admin/source" aria-expanded="true">
                Source <span class="caret"></span>
            </a>
            <ul class="dropdown-menu">
                <li><a href="/admin/source/add">Add</a></li>
                <li><a href="/admin/source/">Manage</a></li>
            </ul>
        </li>
    </ul>
</div>
{{ end }}