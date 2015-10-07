<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/log">Event log</a></li>
      <li class="active">Process #{{ .process.Action.ID }}</li>
    </ul>
  </div>
  <div class="panel-body">
    <div class="table-responsive">
      <span class="text-warning">
				<span class="glyphicon glyphicon-warning-sign"></span>
        This page is not live updated.
			</span>
      <h2>Process #{{ .process.Action.ID }}</h2>
      <div class="row">
        <div class="col-lg-4 col-md-4 col-sm-4">
            <table class="table">
                <tbody>
                    <tr>
                        <td>Type</td>
                        <td><strong>{{ .process.Action.Content }}</strong></td>
                    </tr>
                    <tr>
                        <td>Repository</td>
                        <td><strong>{{ .process.Repository.Name }}</strong></td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div class="col-lg-4 col-md-4 col-sm-4">
            <table class="table">
                <tbody>
                    <tr>
                        <td>Started on:</td>
                        <td>{{ .process.GetStartDate }}</td>
                    </tr>
                    <tr>
                        <td>Finished on:</td>
                        <td><strong>{{ .process.GetEndDate }}</strong></td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div class="col-lg-4 col-md-4 col-sm-4">
            <table class="table">
                <tbody>
                    <tr>
                        <td>Total duration:</td>
                        <td>{{ .process.GetDuration }}</td>
                    </tr>
                    <tr>
                        <td>Current operation:</td>
                        <td><strong>
                          {{ if .operation }}
                          <a href="/admin/log/process/{{.process.Action.ID}}/operation/{{.operation.ID}}">{{ .operation.Content }}</a>
                          {{ else }}
                          -
                          {{ end }}
                        </strong></td>
                    </tr>
                </tbody>
            </table>
        </div>
      </div>
      <table id="list-accounts" class="table table-hover table-bordered table-condensed">
        <thead>
          <tr>
            <th class="hidden-xs">#</th>
            <th>Content</th>
            <th>State</th>
            <th>Start</th>
            <th>End</th>
            <th>Duration</th>
          </tr>
        </thead>
        <tbody>
          {{ $p := .process}}
          {{range $index, $operation := .operations}}
          <tr>
            <td class="col-md-1"><a href="/admin/log/process/{{ $p.Action.ID }}/operation/{{ $operation.ID }}">{{ $operation.ID }}</a></td>
            <td class="col-md-1">{{ $operation.Action.Content }}</td>
            <!-- start state -->
            <td class="col-md-1">
            {{ if $operation.Action.IsRunning }}
            <span class="text-warning">Working</span>
            {{ else }}
            Finished
            {{ end }}
            </td>
            <!-- end state -->
            <td class="col-md-3">{{ $operation.GetStartDate }}</td>
            <td class="col-md-3">{{ $operation.GetEndDate }}</td>
            <td class="col-md-2">{{ $operation.GetDuration }}</td>
          </tr>
          {{end }}
        </tbody>
      </table>
    </div>
  </div>
</div>
