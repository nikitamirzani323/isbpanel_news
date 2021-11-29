<script>
    import { link } from "svelte-spa-router";
    async function handleLogout() {
        localStorage.clear();
        window.location.href = "/";
    }
</script>

<header>
    <nav
        class="navbar navbar-expand-md navbar-dark fixed-top bg-dark"
        style="background-color: #1f2937!important;">
        <div class="container-fluid">
            <a class="navbar-brand" href="/" style="font-size: 13px;">BACKEND</a>
            <button
                class="navbar-toggler"
                type="button"
                data-bs-toggle="collapse"
                data-bs-target="#navbarCollapse"
                aria-controls="navbarCollapse"
                aria-expanded="false"
                aria-label="Toggle navigation">
                <span class="navbar-toggler-icon" />
            </button>
            <div class="collapse navbar-collapse" id="navbarCollapse">
                <ul class="navbar-nav me-auto mb-2 mb-md-0">
                    <li class="nav-item">
                        <a class="nav-link active" aria-current="page" use:link href="/" style="font-size: 13px;">NEWS</a>
                    </li>
                </ul>
                <form class="d-flex">
                    <button 
                        on:click={() => {
                        handleLogout();
                        }}
                        class="btn btn-warning btn-sm">LOGOUT</button>
                </form>
            </div>
        </div>
    </nav>
</header>