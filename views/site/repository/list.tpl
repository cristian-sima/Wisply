
<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li class="active">Repositories</li>
    </ul>
  </div>
  <div class="panel-body">
    <section>
      <h4>Options</h4>
      <div class="btn-group">
        <a href="/admin/repositories/add" class="btn btn-primary">Add repository</a>
      </div>
    </section>
    <br />
    <section>
      <h4>Repositories</h4>
    {{ if .anything }}
    <div class="table-responsive">
      <table class="table table-striped table-hover ">
        <thead>
          <tr>
            <th>Name</th>
            <th>Base URL</th>
            <th>Description</th>
            <th>Modify</th>
            <th>Delete</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $element := .repositories}}
          {{$safe := $element.Name|html}}
          <tr>
            <td>{{ $element.Name |html }}</td>
            <td><a href="{{ $element.URL }}" target="_blank">{{ $element.URL |html }}</a></td>
            <td>{{ $element.Description |html }}</td>
            <td>
              <a href="/admin/repositories/modify/{{$element.ID}}">Modify</a>
            </td>
            <td>
              <a class="deleteRepositoryButton" data-id="{{$element.ID}}" data-name="{{$safe}}" href="/">Delete</a>
            </td>
          </tr>
          {{end }}
        </tbody>
      </table>
    </div>
    {{ else }}
    There are no repositories... :(
    {{ end }}
  </div>
</div>
<script src="/static/js/admin/repository/list.js"></script>
