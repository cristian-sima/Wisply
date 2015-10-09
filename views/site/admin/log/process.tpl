<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/log">Event log</a></li>
      <li class="active">Process #{{ .process.Action.ID }}</li>
    </ul>
  </div>
  <div class="panel-body">
    <div>
      <span class="text-warning">
			<span class="glyphicon glyphicon-warning-sign"></span>
        This page is not live updated.
			</span>
    </div>
    <h2>
      Process #{{ .process.Action.ID }}
    </h2>
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
                    <td><span class="glyphicon glyphicon-calendar"></span> Start:</td>
                    <td>{{ .process.GetStartDate }}</td>
                </tr>
                <tr>
                    <td><span class="glyphicon glyphicon-calendar"></span> Finish:</td>
                    <td><strong>{{ .process.GetEndDate }}</strong></td>
                </tr>
              </tbody>
          </table>
      </div>
      <div class="col-lg-4 col-md-4 col-sm-4">
          <table class="table">
              <tbody>
                  <tr>
                      <td>
                        {{ if .process.Action.IsRunning }}
                        <span class="label label-warning">Working</span>
                        {{ else }}
                        <span class="label label-success">Finished</span>
                        {{ end }}
                      </td>
                      <td>
                        <strong>
                          {{ if .process.Action.IsRunning }}
                          <a href="/admin/log/process/{{.process.Action.ID}}/operation/{{.operation.ID}}">{{ .operation.Content }}</a>
                          {{ else }}
                          <span class="glyphicon glyphicon-time"></span> {{ .process.GetDuration }}
                          {{ end }}
                      </strong>
                    </td>
                  </tr>
              </tbody>
          </table>
      </div>
    </div>
    <div>
      <a href="/admin/log/process/{{ .process.Action.ID }}/history#history">Show history</a> <br />
      <a href="/admin/log/process/{{ .process.Action.ID }}/advance-options">Advance options</a>
    </div>
    <br />
    {{ $len := .operations | len }}
    {{ if eq $len 0 }}
      <div class="text-center">
        There are no operations for this process
      </div>
    {{ else }}
      <div class="table-responsive">
        <table id="list-operations" class="table table-bordered table-condensed">
          <thead>
            <tr>
              <th class="hidden-xs">#</th>
              <th>Content</th>
              <th>State</th>
              <th>Start</th>
              <th>End</th>
              <th><span class="glyphicon glyphicon-time"></span> Duration</th>
            </tr>
          </thead>
          <tbody>
            {{ $p := .process}}
            {{range $index, $operation := .operations}}
            <tr class="{{ $operation.GetResult }}">
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
              <td class="col-md-2">
                {{ if eq $operation.GetDuration "..." }}
                <img src='/static/img/wisply/load.gif' style='height: 20px; width: 20px' />
                {{ else }}
                {{ $operation.GetDuration }}
                {{ end }}
              </td>
            </tr>
            {{end }}
          </tbody>
        </table>
      </div>
    {{ end }}
  </div>
</div>
