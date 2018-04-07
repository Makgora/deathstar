<!DOCTYPE html>
<html>
<head>
{{template "base/head.tpl"}}
    <link rel="stylesheet" href="static/team.css">
</head>

<body>
{{template "base/navbar.tpl"}}

<div class="container mgt">
    <h1 class="display-1 text-center">Qui sommes nous ?</h1>
    <p class="lead text-center" style="margin-top: 50px">
        Vivamus sagittis lacus vel augue laoreet rutrum faucibus dolor auctor. Duis mollis, est non commodo luctus.
    </p>
</div>

<div class="container mgt">
    <div class="row">
        <div class="col-8 offset-2 col-sm-10 offset-sm-1 col-md-8 offset-md-2 col-lg-5 offset-lg-0 col-xl-4 offset-xl-1 mgb">
            <div class="card">
                <img class="card-img-top" src="static/img/team/photo_ilan.jpg" alt="Photo Ilan COHEN" height="400px">
                <div class="card-body">
                    <h4 class="card-title">Ilan COHEN</h4>
                    <p class="card-text text-muted">CEO & Co-founder</p>
                    <p class="font-italic">ilan.cohen@parked.fr</p>
                    <a href="https://www.linkedin.com/in/ilan-chn/" target="_blank">
                        <img src="static/img/logo_social/512px/005-linkedin.png" width="30" />
                    </a>
                    <a href="mailto:titouan.parand@parked.fr">
                        <img src="static/img/logo_social/512px/004-email.png" width="30" />
                    </a>
                </div>
            </div>
        </div>
        <div class="col-8 offset-2 col-sm-10 offset-sm-1 col-md-8 offset-md-2 col-lg-5 offset-lg-2 col-xl-4 offset-xl-2 mgb">
            <div class="card">
                <img class="card-img-top" src="static/img/team/photo_titouan.jpg" alt="Photo Titouan PARAND" height="400px">
                <div class="card-body">
                    <h4 class="card-title">Titouan PARAND</h4>
                    <p class="card-text text-muted">CTO & Co-founder</p>
                    <p class="font-italic">titouan.parand@parked.fr</p>
                    <a href="https://www.linkedin.com/in/titouanparand/" target="_blank">
                        <img src="static/img/logo_social/512px/005-linkedin.png" width="30" />
                    </a>
                    <a href="mailto:titouan.parand@parked.fr">
                        <img src="static/img/logo_social/512px/004-email.png" width="30" />
                    </a>
                    <a href="https://github.com/Makgora" target="_blank">
                        <img src="static/img/logo_social/512px/006-github.png" width="30" />
                    </a>
                </div>
            </div>
        </div>
    </div>
</div>
{{template "base/footer.tpl"}}
{{template "base/scripts.tpl"}}
</body>
</html>