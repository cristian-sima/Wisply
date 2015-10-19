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
          <h2>Wisply tables</h2>
          <br />
          All the tables are generated every day.
          <br />
          Wisply makes available the data which is using as CSV files. In case you need other format, you may want to use <a href="http://www.convertcsv.com/" target="_blank"> this online tool</a>.
          <br />
          <br />
          <br />
          <div class="table-responsive">
            {{ if eq (.tables | len) 0 }}
            There are no tables to be downloaded :(
            {{ else }}
            <table id="list-accounts" class="table table-striped table-hover ">
              <thead>
                <tr>
                  <th class="col-md-1" >#</th>
                  <th class="col-md-4">Table</th>
                  <th class="col-md-5">Description</th>
                  <th class="col-md-2 text-center">Action</th>
                </tr>
              </thead>
              <tbody>
                {{range $index, $table := .tables}}
                <tr>
                  <td class="col-md-1">{{ $table.ID }}</td>
                  <td class="col-md-4">{{ $table.Name }}</td>
                  <td class="col-md-5">{{ $table.Description }}</td>
                  <td class="col-md-2 text-center">
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
