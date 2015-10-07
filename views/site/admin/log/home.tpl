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
      <span class="text-warning">
				<span class="glyphicon glyphicon-warning-sign"></span>
        This page is not live updated.
			</span>
      <table id="list-accounts" class="table table-hover table-bordered table-condensed">
        <thead>
          <tr>
            <th class="hidden-xs">#</th>
            <th>Type</th>
            <th>Repository</th>
            <th>Start</th>
            <th>End</th>
            <!-- <th>Current Operation</th> -->
          </tr>
        </thead>
        <tbody>
          {{range $index, $element := .processes}}
          <tr>
            <td class="col-md-1">{{ $element.ID }}</td>
            <td class="col-md-1">{{ $element.Repository.ID }}</td>
            <td class="col-md-1">{{ $element.Action.Content }}</td>
            <td class="col-md-5">{{ $element.GetStartDate }}</td>
            <td class="col-md-5">{{ $element.GetEndDate }}</td>
          </tr>
          {{end }}
        </tbody>
      </table>
    </div>
  </div>
</div>
