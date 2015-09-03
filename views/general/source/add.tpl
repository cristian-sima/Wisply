<div class="col-lg-8 col-md-7 col-sm-6" >
    <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
            <ul class="breadcrumb">
                <li><a href="/admin">Admin</a></li>
                <li><a href="/admin/source">Sources</a></li>
                <li class="active">Add</li>
            </ul></div>
        <div class="panel-body">
            <p>
            <form action="/admin/source/add" method="POST" class="form-horizontal" >
                <fieldset>
                    <legend>Add a source</legend>
                    <div class="form-group">
                        <label for="sourceName" class="col-lg-2 control-label">Name</label>
                        <div class="col-lg-10">
                            <input type="text" class="form-control" id="sourceName" placeholder="Name">
                        </div>
                    </div>
                    <div class="form-group">
                        <label for="sourceURL" class="col-lg-2 control-label">URL</label>
                        <div class="col-lg-10">
                            <input type="text" class="form-control" id="sourceURL" placeholder="URL">
                        </div>
                    </div>
                    <div class="form-group">
                        <label for="sourceDescription" class="col-lg-2 control-label">Description</label>
                        <div class="col-lg-10">
                            <textarea class="form-control" rows="3" id="sourceDescription"></textarea>
                            <span class="help-block">This field may contain notes or the description of the intitution.</span>
                        </div>
                    </div>
                    
                    <div class="form-group">
                        <div class="col-lg-10 col-lg-offset-2">
                            <button type="reset" class="btn btn-default">Cancel</button>
                            <button type="submit" class="btn btn-primary">Submit</button>
                        </div>
                    </div>
                </fieldset>
            </form>
            </p>
        </div>
    </div>
</div>

<script>
$(document).ready(function() {
    $("#sourceName").focus();
});
</script>