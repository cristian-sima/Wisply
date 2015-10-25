<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li class="active">API</li>
    </ul>
  </div>
  <div class="panel-body">
    <div>
      This panel can be used to allow or reject tables from being downloaded by the public. <Br />
      The tables which are not listed here are by default <strong>rejected</strong>.
      <br />
      <br />
  </div>
  <div class="btn-group">
    <a href="/admin/api/add" class="btn btn-primary"><span class="glyphicon glyphicon-plus"></span> Allow table</a>
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
            <td><a href="/admin/api/modify/{{ $table.ID }}">Modify</a></td>
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
<script src="/static/js/admin/api/list.js"></script>
