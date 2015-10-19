<div class="page-header">
  <div class="row">
    <div class="col-lg-12 col-md-12 col-sm-12">
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/api">API & Developers</a></li>
            <li class="active">List of tables</li>
          </ul>
        </div>
        <div class="panel-body">
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
                  <th>Action</th>
                </tr>
              </thead>
              <tbody>
                {{range $index, $table := .tables}}
                <tr>
                  <td>{{ $table.ID }}</td>
                  <td>{{ $table.Name }}</td>
                  <td>
                      <input type="button" data-name="{{ $table.Name }}" class="btn btn-primary download-table" data-id="{{ $table.ID }}" value="Download" />
                  </td>
                </tr>
                {{end }}
              </tbody>
            </table>
            {{ end }}
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
<script src="/static/js/api/table/list.js"></script>
