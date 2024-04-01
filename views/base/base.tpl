<!DOCTYPE html>
<html lang="ru">
<head>
    {{template "base/head.tpl" .}}
    {{template "head" .}}
</head>
<body class="layout-fluid">
<div class="page">
    {{template "base/navbar.tpl" .}}
    <div class="page-wrapper">
        <div class="page-header d-print-none">
            <div class="container-xl">
                <div class="row g-2 align-items-center">
                    <div class="col">
                        <h2 class="page-title">
                            {{.title}}
                        </h2>
                    </div>
                </div>
            </div>
        </div>
        <div class="page-body">
            {{template "body" .}}
        </div>

    </div>

    {{template "base/footer.tpl" .}}

</div>
<script src="/static/js/tabler.min.js" defer></script>
</body>
</html>