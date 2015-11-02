<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li class="active">Developers &amp; Research</li>
    </ul>
  </div>
  <div class="panel-body">
    <div>
      This panel can be used to allow or reject tables from being downloaded by the public. The public list is available <a target="_blank" href="http://localhost:8081/developer/table/list">here</a>.
      <br />
      The tables which are not listed here are by default <strong>rejected</strong>.
      <br />
      <br />
  </div>
  <div class="btn-group">
    <a href="/admin/developers/add" class="btn btn-primary">
      <span class="glyphicon glyphicon-plus"></span>
      Add table on API public download list
    </a>
  </div>
  <br />
  <br />
    <div class="table-responsive">
      {{ if eq (.tables | len) 0 }}
      There are no tables to be downloaded :(
      {{ else }}
      <table id="list-accounts" class="table table-striped table-hover ">
        <thead>
          <tr>
            <th class="hidden-xs">#</th>
            <th>Table</th>
            <th>Modify</th>
            <th>Private</th>
          </tr>
        </thead>
        <tbody>
          {{ $xsrf_input := .xsrf_input }}
          {{range $index, $table := .tables}}
          <tr>
            <td>{{ $table.ID }}</td>
            <td>{{ $table.Name }}</td>
            <td><a href="/admin/developers/table/{{ $table.ID }}/modify">Modify</a></td>
            <td>
              <a href="#" class="makeTablePrivate" data-id="{{ $table.ID }}" data-name="{{ $table.Name }}">Make private</a>
            </td>
          </tr>
          {{end }}
        </tbody>
      </table>
      {{ end }}
    </div>
  </div>
</div>
<script src="/static/js/admin/developers/home.js"></script>
<script>
$(document).ready(function(){
    var module = wisply.getModule("admin-developers-home"),
      list = new module.List();
      list.init();
});
</script>
