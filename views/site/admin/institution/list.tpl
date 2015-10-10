
<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li class="active">Institutions</li>
    </ul>
  </div>
  <div class="panel-body">
    <section>
      <h4>Options</h4>
      <div class="btn-group">
        <a href="/admin/institutions/add" class="btn btn-primary">Add institution</a>
      </div>
    </section>
    <br />
    <section>
      <div class="row">
        <div class="col-md-1">
          <h4>Institutions</h4>
        </div>
        <div class="col-md-11 text-right">
          <div id="websocket-connection"></div>
        </div>
      </div>
    {{ if .anything }}
    <div class="table-responsive">
      <table class="table table-striped table-hover " id="institution-list">
        <thead>
          <tr>
            <th>Name</th>
            <th>Web adddress</th>
            <th>Auto Wiki</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $element := .institutions}}
          {{$safe := $element.Name|html}}
          <tr>
            <td><a href="/admin/institutions/institution/{{ $element.ID }}">{{ $element.Name |html }}</a></td>
            <td><a href="{{ $element.URL }}" target="_blank">{{ $element.URL |html }}</a></td>
            <td>
            {{ if eq $element.WikiID "NULL" }}
            <span class="label label-danger">Disabled</label>
            {{ else }}
            <span class="label label-success">Enabled</label>
            {{ end }}
            </td>
          </tr>
          {{end }}
        </tbody>
      </table>
    </div>
    {{ else }}
    There are no institution... :(
    {{ end }}
  </div>
</div>
<script src="/static/js/admin/institution/list.js"></script>
