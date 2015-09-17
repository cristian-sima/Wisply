
<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li class="active">Sources</li>
    </ul>
  </div>
  <div class="panel-body">
    <section>
      <h4>Options</h4>
      <div class="btn-group">
        <a href="/admin/sources/add" class="btn btn-primary">Add source</a>
      </div>
    </section>
    <br />
    <section>
      <h4>Sources</h4>
    {{ if .anything }}
    <div class="table-responsive">
      <table class="table table-striped table-hover ">
        <thead>
          <tr>
            <th>Name</th>
            <th>URL</th>
            <th>Description</th>
            <th>Modify</th>
            <th>Delete</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $element := .sources}}
          {{$safe := $element.Name|html}}
          <tr>
            <td>{{ $element.Name |html }}</td>
            <td><a href="{{ $element.URL }}" target="_blank">{{ $element.URL |html }}</a></td>
            <td>{{ $element.Description |html }}</td>
            <td>
              <a href="/admin/sources/modify/{{$element.ID}}">Modify</a>
            </td>
            <td>
              <a class="deleteSourceButton" data-id="{{$element.ID}}" data-name="{{$safe}}" href="/">Delete</a>
            </td>
          </tr>
          {{end }}
        </tbody>
      </table>
    </div>
    {{ else }}
    There are no sources... :(
    {{ end }}
  </div>
</div>
<script src="/static/js/admin/source/list.js"></script>
