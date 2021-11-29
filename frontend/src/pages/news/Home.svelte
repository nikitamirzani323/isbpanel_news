<script>
    import { Input } from "sveltestrap";
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";
    export let table_header_font = ""
	export let table_body_font = ""
	export let token = ""
	export let listHome = []
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
    let title_page = "NEWS"
    let sData = "";
    let myModal = "";
    
    let listnews = []
    let listcategory = []
    let record = ""
    let totalrecordnews = 0
    let totalrecordcategory = 0
    
    let tanggal_start_newsfetch = "";
    let tanggal_end_newsfetch = "";
    let page_newsfetch = "";
    let category_field_idrecord = 0;
    let category_field_name = "";
    let category_field_display = 0;
    let category_field_status = "";
    let news_field_idrecord = 0;
    let news_field_title = "";
    let news_field_descp = "";
    let news_field_category = "";
    let news_field_url = "";
    let news_field_image = "";
    let searchNews = "";
    let filterNews = "";
    let css_loader = "display: none;";
    let msgloader = "";
    $: {
        if (searchNews) {
            filterNews = listHome.filter(
                (item) =>
                    item.news_title
                        .toLowerCase()
                        .includes(searchNews.toLowerCase())
            );
        } else {
            filterNews = [...listHome];
        }
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const ShowFormNewsFetch = () => {
        sData = "Edit"
        myModal = new bootstrap.Modal(document.getElementById("modalfetchnew"));
        myModal.show();
    };
    
    
    const ShowFormNews = (e,id,category,title,descp,url,image) => {
        sData = e
        news_field_idrecord = parseInt(id);
        news_field_title = title;
        news_field_descp = descp;
        news_field_category = parseInt(category);
        news_field_url = url;
        news_field_image = image;
        call_category()
        myModal = new bootstrap.Modal(document.getElementById("modalcrudnews"));
        myModal.show();
    };
    async function call_news() {
        listnews = [];
        let KEY_NEWS = "apiKey=25ff185c903e49ddba06551850241e06"
        let COUNTRY_NEWS = "country=id"
        let PAGE_NEWS = "page="+page_newsfetch
        let FROM_NEWS = "from=" + tanggal_start_newsfetch
        let TO_NEWS = "to="+tanggal_end_newsfetch
        let URL_NEWS = "https://newsapi.org/v2/top-headlines?"+KEY_NEWS+"&"+COUNTRY_NEWS+"&"+FROM_NEWS+"&"+TO_NEWS+"&"+PAGE_NEWS
        const res = await fetch(URL_NEWS);
        const json = await res.json();
        let status = json.status;
        let message = json.message;
        let record = json.articles;
        let no = 0;
        if(status == "ok"){
            totalrecordnews = record.length;
            for (var i = 0; i < record.length; i++) {
                no = no + 1
                listnews = [
                            ...listnews,
                    {
                        news_no: no,
                        news_author: record[i]["author"],
                        news_title: record[i]["title"],
                        news_description: record[i]["description"],
                        news_url: record[i]["url"],
                        news_urlToImage: record[i]["urlToImage"],
                        news_publishedat: record[i]["publishedAt"],
                        news_content: record[i]["content"],
                    },
                ];
            }
        }else{
            alert(message)
        }
        
       
        
    }
    async function call_category() {
        listcategory = [];
        const res = await fetch("/api/categorynews", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            if (record != null) {
                totalrecordcategory = record.length;
                let no = 0
                for (var i = 0; i < record.length; i++) {
                    if(parseInt(record[i]["category_id"]) != 2112){
                        no = no + 1;
                        listcategory = [
                        ...listcategory,
                            {
                                category_no: no,
                                category_id: record[i]["category_id"],
                                category_name: record[i]["category_name"],
                                category_totalnews: record[i]["category_totalnews"],
                                category_display: record[i]["category_display"],
                                category_status: record[i]["category_status"],
                                category_statuscss: record[i]["category_statuscss"],
                                category_create: record[i]["category_create"],
                                category_update: record[i]["category_update"],
                            },
                        ];
                    }
                }
            }
        } 
    }
    
    async function handleSave() {
        let flag = true
        let msg = ""
        css_loader = "display: inline-block;";
        msgloader = "Sending...";
        const res = await fetch("/api/newssave", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sdata: sData,
                page:"NEWS-SAVE",
                news_id: news_field_idrecord,
                news_category: news_field_category,
                news_title: news_field_title,
                news_descp: news_field_descp,
                news_url: news_field_url,
                news_image: news_field_image,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            msgloader = json.message;
            myModal.hide()
            RefreshHalaman()
        } else if(json.status == 403){
            alert(json.message)
        } else {
            msgloader = json.message;
        }
        setTimeout(function () {
            css_loader = "display: none;";
        }, 1000);
    }
    async function handleDeleteNews(e) {
        let flag = true
        let msg = ""
        if(e == ""){
            flag = false
            msg = "The News is required"
        }
        if(flag){
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/newsdelete", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page:"NEWS-DELETE",
                    news_id: parseInt(e),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                RefreshHalaman()
                msgloader = json.message;
            } else if(json.status == 403){
                alert(json.message)
            } else {
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
   
    function callFunction(event){
        switch(event.detail){
            case "CALL_FORMNEWS":
                ShowFormNewsFetch();
                break;
            case "FETCH_NEWS":
                call_news();
                break;
        
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE_NEWS":
                handleSave();break;
        }
    }
    function clearfield_category(){
        category_field_idrecord = 0;
        category_field_name = "";
        category_field_display = 0;
        category_field_status = "";
    }
    const handleKeyboard_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
                filterNews = [];
                listHome = [];
                const news = {
                    searchNews,
                };
                dispatch("handleNews", news);
        }  
    };
</script>

<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container-fluid" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-6">   
            <Button
                on:click={callFunction}
                button_function="CALL_FORMNEWS"
                button_title="Fetch"
                button_css="btn-primary"/>
            <Panel
                card_title="{title_page}"
                card_footer={totalrecordnews}>
                <slot:template slot="card-body">
                        <table class="table table-striped table-hover">
                            <thead>
                                <tr>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;">&nbsp;</th>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NEWS</th>
                                </tr>
                            </thead>
                            {#if totalrecordnews > 0}
                            <tbody>
                                {#each listnews as rec }
                                    <tr>
                                        <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                            <i 
                                                on:click={() => {
                                                    ShowFormNews("New",0,"",rec.news_title,rec.news_description,rec.news_url,rec.news_urlToImage);
                                                }} 
                                                class="bi bi-save"></i>
                                        </td>
                                        
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.news_no}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                            <b>AUTHOR : </b> {rec.news_author}<br>
                                            <a href="{rec.news_url}" target="_blank">{rec.news_title}</a><br>
                                            {@html rec.news_description} <br>
                                            <img width="100" src="{rec.news_urlToImage}" class="img-thumbnail" alt="...">
                                        </td>
                                    </tr>
                                {/each}
                            </tbody>
                            {:else}
                            <tbody>
                                <tr>
                                    <td colspan="20">
                                        <center>
                                            <Loader />
                                        </center>
                                    </td>
                                </tr>
                            </tbody>
                            {/if} 
                        </table>
                </slot:template>
            </Panel>
        </div>
        <div class="col-sm-6">
            <Button
                on:click={callFunction}
                button_function="REFRESH"
                button_title="Refresh"
                button_css="btn-primary"/>
            
            <Panel
                card_search={true}
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-search">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchNews}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search News"
                            aria-label="Search"
                        />
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                        <table class="table table-striped table-hover table-sm">
                            <thead>
                                <tr>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan="2">&nbsp;</th>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">NEWS</th>
                                </tr>
                            </thead>
                            {#if totalrecord > 0}
                            <tbody>
                                {#each filterNews as rec }
                                    <tr>
                                        <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                            <i 
                                                on:click={() => {
                                                    ShowFormNews("Edit",rec.news_id,rec.news_idcategory,rec.news_title,rec.news_descp,rec.news_url,rec.news_image)
                                                }} 
                                                class="bi bi-pencil"></i>
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                            <i 
                                                on:click={() => {
                                                    handleDeleteNews(rec.news_id);
                                                }} 
                                                class="bi bi-trash"></i>
                                        </td>
                                        <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.news_no}</td>
                                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                            CATEGORY : {rec.news_category}<br>
                                            <a href="{rec.news_url}" target="_blank">{rec.news_title}</a><br>
                                            <img width="100" src="{rec.news_image}" class="img-thumbnail" alt="...">
                                        </td>
                                    </tr>
                                {/each}
                            </tbody>
                            {:else}
                            <tbody>
                                <tr>
                                    <td colspan="20">
                                        <center>
                                            <Loader />
                                        </center>
                                    </td>
                                </tr>
                            </tbody>
                            {/if} 
                        </table>
                </slot:template>
            </Panel>
        </div>
    </div>
</div>
<Modal
	modal_id="modalfetchnew"
	modal_size="modal-dialog-centered"
	modal_title="{title_page}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="example" class="form-label">Start</label>
			<Input
                bind:value={tanggal_start_newsfetch}
                class="required"
                type="date"
                name="date"
                id="exampleDate"
                data-date-format="dd-mm-yyyy"
                placeholder="date placeholder"/>
		</div>
        <div class="mb-3">
            <label for="example" class="form-label">End</label>
			<Input
                bind:value={tanggal_end_newsfetch}
                class="required"
                type="date"
                name="date"
                id="exampleDate"
                data-date-format="dd-mm-yyyy"
                placeholder="date placeholder"/>
		</div>
        <div class="mb-3">
            <label for="example" class="form-label">Page</label>
            <select 
                class="form-control"
                bind:value={page_newsfetch}>
                <option value="1">Page 1</option>
                <option value="2">Page 2</option>
                <option value="3">Page 3</option>
            </select>
		</div>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="FETCH_NEWS"
            button_title="Submit"
            button_css="btn-warning"/>
	</slot:template>
</Modal>
<Modal
	modal_id="modalcrudnews"
	modal_size="modal-dialog-centered"
	modal_title="NEWS/{sData}"
    modal_body_css=""
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Category</label>
            <select 
                bind:value={news_field_category}
                class="form-control required">
                {#each listcategory as rec}
                <option value="{rec.category_id}">{rec.category_name}</option>
                {/each}
            </select>
		</div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Title</label>
			<Input
                bind:value={news_field_title}
                class="required"
                type="text"
                placeholder="News Title"/>
		</div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Deskripsi</label>
            <textarea
                style="height: 100px;resize: none;" 
                bind:value={news_field_descp} class="form-control required"></textarea>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Url Image</label>
			<Input
                bind:value={news_field_image}
                class="required"
                type="text"
                placeholder="News Image"/>
		</div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Url</label>
			<Input
                bind:value={news_field_url}
                class="required"
                type="text"
                placeholder="News URL"/>
		</div>
	</slot:template>
	<slot:template slot="footer">
        <Button
            on:click={callFunction}
            button_function="SAVE_NEWS"
            button_title="Save"
            button_css="btn-warning"/>
	</slot:template>
</Modal>
