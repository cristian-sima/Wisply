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
    <a href="/admin/api/add" class="btn btn-primary">Add table</a>
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
            <th>Delete</th>
            <th>Modify</th>
          </tr>
        </thead>
        <tbody>
          {{ $xsrf_input := .xsrf_input }}
          {{range $index, $table := .tables}}
          <tr>
            <td>{{ $table.ID }}</td>
            <td>{{ $table.Name }}</td>
            <td>
              <form action="/admin/api/delete" method="POST">
          			{{ $xsrf_input }}
                <input type="hidden" name="table-id" value="{{ $table.ID }}"/>
                <input type="submit" class="btn btn-primary" value="Remove" />
              </form>
            </td>
            <td><a href="/admin/api/modify/{{ $table.ID }}">Modify</a></td>
          </tr>
          {{end }}
        </tbody>
      </table>
      {{ end }}
    </div>
  </div>
</div>
<script src="/static/js/admin/account/list.js"></script>
