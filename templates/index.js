const resultTable = document.getElementById("resultTable");
const title = document.getElementById("result-title");
const noResultsText = document.getElementById("no-results-text");

document.getElementById("search").onsubmit = async ev => {
    ev.preventDefault();
    const lat = ev.target["lat"].value;
    const lon = ev.target["lon"].value;
    const r = ev.target["r"].value;

    if (lat && lon && r) {
        const result = await fetch(`/list?lat=${lat}&lon=${lon}&r=${r}`);
        const resultJson = await result.json();
        title.style.visibility = "visible";
        if (resultJson.length > 0) {
            resultTable.innerHTML = `<tr><th>Name</th><th>Distance (km)</th><th>Latitude</th><th>Longitude</th></tr>`;
            resultTable.innerHTML += resultJson.map(item => {
                const {distance, name, coords: {lat, lon}} = item;
                return `<tr><th>${name}</th><th>${distance}</th><th>${lat}</th><th>${lon}</th></tr>`;
            }).join("");
            resultTable.style.visibility = "visible";
            noResultsText.style.visibility = "hidden";
        } else {
            resultTable.style.visibility = "hidden";
            noResultsText.style.visibility = "visible";
        }
    }
};
