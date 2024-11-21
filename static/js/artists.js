document.addEventListener("DOMContentLoaded", async () => {
    const artistContainer = document.getElementById("artistContainer");

    try {
        // Get artist data from the embedded JSON script tag asynchronously
        const artistDataScript = document.getElementById("artistData");
        const artists = JSON.parse(artistDataScript.textContent);

        // Simulating an asynchronous operation if needed (e.g., processing)
        await new Promise(resolve => setTimeout(resolve, 100)); // Example: artificial delay
        console.log(artists);

        // Render artists in the container
        artistContainer.innerHTML = artists
            .map(
                (artist) => `
                <div class="card">
                    <a href="/artist/${artist.id}">
                        <img src="${artist.image}" alt="${artist.name}">
                        <h2>${artist.name}</h2>
                    </a>
                </div>
            `
            )
            .join("");
    } catch (error) {
        console.error("Failed to process artist data:", error);
        artistContainer.innerHTML = `<p class="error">Failed to load artist data. Please try again later.</p>`;
    }
});
