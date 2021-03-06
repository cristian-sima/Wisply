<div class="page-header">
  <div class="row">
    <div class="col-lg-12 col-md-12 col-sm-12">
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/developer">Developers & Research</a></li>
            <li class="active">Download tables</li>
          </ul>
        </div>
        <div class="panel-body">
          <h1>Wisply tables</h1>
          <div class="well well-sm">
          These tables are updated daily.
          <br />
          Wisply makes the tables available in CSV format.
          <br />
          In case you need them in other format, you may want to use <a href="http://www.convertcsv.com/" target="_blank"> this online tool</a>.
          <br />
        </div>
          <div>
            <span class="text-warning">
              <span class="glyphicon glyphicon-star"></span>
              <span class="glyphicon glyphicon-star"></span>
              <span class="glyphicon glyphicon-star"></span>
              <span class="star-muted">
                <span class="glyphicon glyphicon-star"></span>
                <span class="glyphicon glyphicon-star"></span>
              </span>
            </span>
            &nbsp;&nbsp;&nbsp;&nbsp;
            <a data-toggle="tooltip" data-popover="true" data-content="Tim Berners-Lee, the inventor of the Web and Linked Data initiator, suggested a 5-star deployment scheme for Open Data. The scheme measures how well data is integrated into the web. <a href='http://5stardata.info/en/'>Read more</a>" data-html=true>
            <span class="glyphicon glyphicon-info-sign"></span>
          </a>
          </div>
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
                    <span class="h6 text-muted">{{ $table.Name }}</span>
                    <br />
                    <br />
                    <button id="download-{{$table.Name}}-table" data-name="{{ $table.Name }}" data-id="{{ $table.ID }}" type="button" class="btn btn-primary download-table" aria-label="Left Align">
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
<div>
<style scoped>
.table-name {
  text-transform: capitalize;
}
.star-muted {
  color: #ECECEC;
}
</style>
</div>
<script src="/static/js/developer/data/table.js"></script>
<script>
$(document).ready(function(){
    var module = wisply.getModule("developer-data-table"),
      manager = new module.Manager();
      manager.init();
});
</script>
