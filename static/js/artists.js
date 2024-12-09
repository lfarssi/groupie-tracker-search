const artistContainer = document.getElementById("artistContainer");
const suggestion = document.getElementById("suggestions");
// Get artist data from the embedded JSON script tag asynchronously
const artistDataScript = document.getElementById("artistData");
const artists = JSON.parse(artistDataScript.textContent);
const uniqueLocations = new Set();
const uniqueMember = new Set();
const uniqueCreationDate = new Set();
const uniqueFirstAlbum = new Set();
const uniqueDates = new Set();
console.log(artists);

artists.forEach((artist) => {
  const artistName = document.createElement("option");
  artistName.value = artist.name;
  artistName.innerHTML = "- Artist/Band";
  suggestion.appendChild(artistName);

  if (!uniqueFirstAlbum.has(artist.firstAlbum)) {
    uniqueFirstAlbum.add(artist.firstAlbum);
    const firstAlbum = document.createElement("option");
    firstAlbum.value = artist.firstAlbum;
    firstAlbum.innerHTML = "- First Album";
    suggestion.appendChild(firstAlbum);
  }
  if (!uniqueCreationDate.has(artist.creationDate)) {
    uniqueCreationDate.add(artist.creationDate);
    const creationDate = document.createElement("option");
    creationDate.value = artist.creationDate;
    creationDate.innerHTML = "- Creation Date";
    suggestion.appendChild(creationDate);
  }

  artist.members.forEach((member) => {
    if (!uniqueMember.has(member)) {
      uniqueMember.add(member);
    const members = document.createElement("option");
    members.value = member;
    members.innerHTML = "- Member";
    suggestion.appendChild(members);
    }
  });

  artist.Locationsr.forEach((city) => {
    if (!uniqueLocations.has(city)) {
      uniqueLocations.add(city);
      const locations = document.createElement("option");
      locations.value = city;
      locations.innerHTML = "- Location";
      suggestion.appendChild(locations);
    }
  });

  artist.ConcertDatesr.forEach((date) => {
    if (!uniqueDates.has(date)) {
      uniqueDates.add(date);
      const dates = document.createElement("option");
      dates.value = date;
      dates.innerHTML = "- Concert Date";
      suggestion.appendChild(dates);
    }
  });
});

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
  )
  .join("");
if (artists.length == 0) {
  artistContainer.innerHTML = `
    <div class="card">
        <h2>No artists found</h2>
    </div>
    `;
}
