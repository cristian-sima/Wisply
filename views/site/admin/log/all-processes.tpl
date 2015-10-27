<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li class="active">Event log</li>
    </ul>
  </div>
  <div class="panel-body">
    <br />
    {{ $len := .processes | len }}
    {{ if eq $len 0 }}
    <div class="text-center">
      There is no recorded process
    </div>
    {{ else }}
    <div>
      <span class="text-warning warning-notice">
        <span class="glyphicon glyphicon-warning-sign"></span>
        This page is not live updated.
      </span>
      <br />
    </div>
    <div class="table-responsive">
      <table id="list-processes" class="table table-bordered table-condensed">
        <thead>
          <tr>
            <th class="hidden-xs">#</th>
            <th>Type</th>
            <th><span class="glyphicon glyphicon-list-alt"></span></th>
            <th>State</th>
            <th>Start</th>
            <th>End</th>
            <th>Duration</th>
            <th>Now</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $element := .processes}}
          <tr class="{{ $element.GetResult }}">
            <td class="col-md-1"><a class="btn btn-primary btn-xs" href="/admin/log/process/{{ $element.Action.ID }}"><span class="no-print">See</span> #{{ $element.Action.ID }}</a></td>
            <td class="col-md-1.5">{{ $element.Action.Content }}</td>
            <td class="col-md-0.5"><a data-toggle="tooltip" title="See progress history" href="/admin/log/process/{{ $element.Action.ID }}/history#history"><span class="glyphicon glyphicon-list-alt"></span></a></td>
            <!-- start state -->
            <td class="col-md-1">
              {{ if $element.IsSuspended }}
              <a href="/admin/log/process/{{ $element.Action.ID }}"><span class="label label-warning">Suspended</span></a>
              {{ else }}
              {{ if $element.Action.IsRunning }}
              <span class="label label-info">Working</span>
              {{ else }}
              <span class="label label-success">Finished</span>
              {{ end }}
              {{ end }}
            </td>
            <!-- end state -->
            <td class="col-md-1.5">{{ $element.GetStartDate }}</td>
            <td class="col-md-1.5">{{ $element.GetEndDate }}</td>
            <td class="col-md-3">
              {{ if eq $element.GetDuration "..." }}
              <img src='/static/img/wisply/load.gif' style='height: 20px; width: 20px' />
              {{ else }}
              {{ $element.GetDuration }}
              {{ end }}
            </td>
            <td class="col-md-1">
              {{ if $element.GetCurrentOperation }}
              <a href="/admin/log/process/{{$element.ID}}/operation/{{$element.GetCurrentOperation.ID}}">{{ $element.GetCurrentOperation.Content }}</a>
              {{ else }}
              -
              {{ end }}
            </td>
          </tr>
          {{end }}
        </tbody>
      </table>
    </div>
    {{ end }}
    <div id="other-options">
      <a class="btn btn-primary" href="/admin/log/advance-options">Advance options</a>
    </div>
  </div>
</div>
<script>
$(document).ready(function() {
  wisply.activateTooltip();
});
</script>
