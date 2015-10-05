<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/repositories">Repositories</a></li>
      <li class="active">Event log</li>
    </ul>
  </div>
  <div class="panel-body">
    <div class="table-responsive">
      <table id="list-accounts" class="table table-hover table-bordered table-condensed">
        <thead>
          <tr>
            <th class="hidden-xs">#</th>
            <th>Date</th>
            <th>Duration</th>
            <th>Repository</th>
            <th>Content</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $element := .events}}
          <tr>
            <td class="col-md-1">{{ $element.ID |html }}</td>
            <td class="col-md-2">{{ $element.Timestamp |html }}</td>
            <td class="col-md-1">{{ $element.Duration }}</td>
            <td class="col-md-1">
                  {{ $element.Repository }}
            </td>
            <td class="col-md-7">
                  {{ $element.Content }}
            </td>
          </tr>
          {{end }}
        </tbody>
      </table>
    </div>
  </div>
</div>
<script src="/static/js/admin/account/list.js"></script>
