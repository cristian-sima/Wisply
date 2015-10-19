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
                  <th class="col-md-1">#</th>
                  <th class="col-md-4">Table</th>
                  <th class="col-md-7">Explication</th>
                </tr>
              </thead>
              <tbody>
                {{range $index, $table := .tables}}
                <tr>
                  <td class="col-md-1">{{ $table.ID }} </td>
                  <td class="col-md-4">
                    <h2 class="table-name" >{{ $table.Name }}</h2>
                    <h5 class="text-muted">{{ $table.Name }}</h5>
                    <br />
                    <button data-name="{{ $table.Name }}" data-id="{{ $table.ID }}" type="button" class="btn btn-primary download-table" aria-label="Left Align">
                      <span class="glyphicon glyphicon-download-alt" aria-hidden="true"></span> Download
                    </button>
                    <br />
                    <br />
                  </td>
                  <td class="col-md-7">{{ $table.GetDescription }}</td>
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
<style>
.table-name {
  text-transform: capitalize;
}
</style>
<script src="/static/js/api/table/list.js"></script>
