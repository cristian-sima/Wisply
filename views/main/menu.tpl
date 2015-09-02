{{ define "menu" }}
<div class="navbar navbar-default navbar-fixed-top">
    <div class="container">
        <div class="navbar-header">
            <a href="/" class="navbar-brand">Wisply</a>
            <button class="navbar-toggle" type="button" data-toggle="collapse" data-target="#navbar-main">
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
        </div>
        <div class="navbar-collapse collapse" id="navbar-main">
            <ul class="nav navbar-nav">
                <li class="dropdown">
                    <a href="http://bootswatch.com/paper/#" id="themes">About</a>
                </li>
                <li>
                    <a href="/webscience">Web science</a>
                </li>
                <li>
                    <a href="/contact">Contact</a>
                </li>                       

            </ul>

            <ul class="nav navbar-nav navbar-right">
                <li><form class="navbar-form navbar-left" role="search">
                        <div class="form-group">
                            <input type="text" class="form-control" placeholder="Search">
                        </div>      
                        <a href="http://bootswatch.com/paper/#" class="btn btn-default">Log in</a>
                        <a href="http://bootswatch.com/paper/#" class="btn btn-default">Sign up</a>
                    </form>
                </li>
            </ul>
        </div>
    </div>
</div>
{{ end }}