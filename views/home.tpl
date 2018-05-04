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
        <h2 class="col-12">
            Se garer devient un jeu d'enfant !
        </h2>
        <p class="col-12">
            Parked vous aide à trouver la place de parking idéale pour vos déplacements.
            <br/>
            Rapide, simple et gratuit !
        </p>
    </div>
    <div class="row">
        <img src="static/img/streetView.jpg">
    </div>
</div>

<div id="how" class="container-fluid mgt">
    <div class="row text-center">
        <h2 class="col-12">
            Garez-vous rapidement et simplement
        </h2>
        <p class="col-12">
            Ne tournez plus dans les rues, une place vous attend !
        </p>
    </div>
    <div class="row">
        <div class="card-group col-10 offset-1">
            <div class="card col-12 col-md-4">
                <img class="card-img-top" src="static/img/Illu_01.jpg" alt="Trouvez votre parking">
                <div class="card-body">
                    <h3>
                        Trouvez le parking idéal
                    </h3>
                    <p class="text-center">
                        Parked vous recommande le parking disponible le plus adapté à votre destination.
                        Partout dans Monaco, une place vous attend.
                    </p>
                </div>
            </div>
            <div class="card col-12 col-md-4">
                <img class="card-img-top" src="static/img/Illu_02.jpg" alt="Préparer votre déplacement">
                <div class="card-body">
                    <h3>
                        Préparez votre déplacement
                    </h3>
                    <p class="text-center">
                        Trouvez toutes les informations utiles sur le parking sélectionné, ses tarifs et les services disponibles.
                    </p>
                </div>
            </div>
            <div class="card col-12 col-md-4">
                <img class="card-img-top" src="static/img/Illu_03.jpg" alt="Garez vous !">
                <div class="card-body">
                    <h3>
                        Garez-vous !
                    </h3>
                    <p class="text-center">
                        Laissez-vous guider jusqu’à votre parking et garez-vous.
                        Partez serein, votre voiture est en sécurité.
                    </p>
                </div>
            </div>
        </div>
    </div>
</div>

<div class="container">
    <div class="row">
        <a href="#" class="btn btn-secondary col-4 offset-4 mgb mgt" role="button">Reserver une place !</a>
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