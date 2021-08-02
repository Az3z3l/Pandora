async function api(operationName, variables, query) {
    var nano = JSON.stringify({
        "operationName": operationName,
        "variables": variables,
        "query": query,
    });
    var response = await fetch("/junior/query", {
        headers: {
            "content-type": "application/json",
        },
        "body": nano,
        "method": "POST",
    });

    if (response.ok){
        let x = await response.json();
        return (x.data)
    } else {
        localStorage.removeItem('user');
        document.location='/login'
    }
}

async function isset(operationName, variables, query) {
   
    var response = await fetch("/api/isset", {
        credentials: "include",
        "method": "POST",
        "mode": "cors"
    });

    if (response.ok){
        let x = await response.json();
        if (x.Status==='ok'){
        localStorage.setItem("Jedi", "May the force be with you");
        document.location='/challenges'
        }
    }
}

export {
    api,
    isset,
}
  

// exports.api=api();
// exports.isset=isset();