<!DOCTYPE html>
<html>
<head>
{{template "base/head.tpl"}}
    <link rel="stylesheet" href="static/css/home.css">
</head>

<body>
{{template "base/navbar.tpl"}}

<div class="container-fluid">
    <div class="row">
        <div id="parkings" class="col-12 col-lg-6">
            {{range .parkings}}
            <div class="parking container">
                <div class="row">
                    <div class="col-2" style="background-color: #dddddd">
                        <img width="100%" class="center" src="static/img/parked/logo_parked.png" alt="Parking photo">
                    </div>
                    <div class="col-10">
                        <div class="row">
                            <div class="col-9">
                                <label>{{.Name}}</label>
                            </div>
                            <div class="col-3">
                                <label>{{.Statut}}</label>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-9">
                                <label class="text-muted">{{.Address}}</label>
                            </div>
                            <div class="col-3">
                                <label>{{.FreeSpaces}} places</label>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-9">
                                <button type="button" class="btn btn-outline-secondary"disabled>Réserver</button>
                            </div>
                            <div class="col-3">
                                {{if not .Address}}
                                    <button type="button" class="btn btn-outline-secondary"disabled>Go</button>
                                {{else}}
                                    <a href="https://www.google.fr/maps/dir//{{.Address}}" target="_blank" type="button" class="btn btn-outline-success btn-success">Go</a>
                                {{end}}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            {{end}}
        </div>
        <div class="d-none d-lg-block col-lg-6">
            <div id="map">
            <script>
                var cityLocation = {lat: {{.cityLocation.Lat}}, lng: {{.cityLocation.Lng}}};
                var beaches = [];
                {{range .parkings}}
                        beaches.push(["{{.Name}}", {{.Location.Lat}}, {{.Location.Lng}}]);
                {{end}}
            </script>
            </div>
        </div>
    </div>
    <div class="row">
        <div class="col-12 col-lg-6">
            <label id="update" class="text-muted font-italic">
                Dernière mise à jour à {{.lastUpdate}}
            </label>
        </div>
    </div>
</div>

<div class="container-fluid mgt">
    <div class="row text-center">
        <div id="streetViewTitle" class="col-12">
            Se garer devient un jeu d'enfant !
        </div>
    </div>
    <div class="row text-center">
        <div id="streetViewDescription" class="col-12">
            Parked vous aide à trouver la place de parking idéale pour vos déplacements.
            Rapide, simple et gratuit !
        </div>
    </div>
    <div class="row">
        <img src="static/img/streetView.jpg">
    </div>
</div>

{{template "base/newsletter.tpl"}}

{{template "base/footer.tpl"}}

{{template "base/scripts.tpl"}}

<script src="https://maps.googleapis.com/maps/api/js?key=AIzaSyBVFpJA9GjD8dqEu0KHVwNjQTZ6XwrnCFQ"></script>
<script src="static/js/map.js"></script>
<script src="static/js/home.js"></script>
</body>

</html>