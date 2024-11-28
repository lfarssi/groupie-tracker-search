    const artistContainer = document.getElementById("artistContainer");


        // Get artist data from the embedded JSON script tag asynchronously
        const artistDataScript = document.getElementById("artistData");
        const artists = JSON.parse(artistDataScript.textContent);


        // Render artists in the container
        artistContainer.innerHTML = artists
            .map(
                (artist) => `
                <div class="card">
                    <a href="/artist/${artist.id}">
                    <div class="div-image">
                        <img src="${artist.image}" alt="${artist.name}">
                    </div>
                        <h2>${artist.name}  </h2>
                        <p>${artist.Type}</p>
                    </a>
                </div>
            `
            ).join("")
   
