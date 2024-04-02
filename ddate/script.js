document.addEventListener("DOMContentLoaded", function() {
    fetch("http://localhost:8080/")
    .then(response => response.json())
    .then(data => {
        const { day, number, season, year, holyday } = data;
        let discordianDate = `${day} of ${season}, YOLD ${year}`;
        
        // Check if holyday exists and append celebrating message
        if (holyday) {
            discordianDate += `<br><span class="celebrating">celebrating ${holyday}</span>`;
        }
        
        document.getElementById("discordian-date").innerHTML = discordianDate;
    })
    .catch(error => {
        console.error("Error fetching Discordian date:", error);
    });
});
