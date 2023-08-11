const graphqlUrl = "http://localhost:2003/v1/graphql"; // Replace with your actual GraphQL endpoint URL

const base64str = "iVBORw0KGgoAAAANSUhEUgAAAIIAAACCCAMAAAC93eDPAAAAY1BMVEX///8AAABAQEDi4uJmZmZ6enrn5+fJyckHBwf29vb5+fnd3d0vLy8kJCQhISFLS0udnZ06OjrPz8+9vb2CgoKIiIipqalXV1cpKSmzs7Pu7u5ra2teXl4RERHV1dWRkZEaGhrM4Hq+AAADGUlEQVR4nO2Z63aqQAyFmUFBYEC5CAKKvP9THjJUUctwCzhnreb72RTYpJtMkhoGQRAEQRAEQRAEQRB/gcJC4+IU5AE7IGGlg1HgsDWI9UtIMBKM+20FCUeUBMOzUUT3RkKAMgOaCtJQaZVQwF8y1yrBzRoJZ60SjDPej1hyMAOyQCLxQIKnVYIBElK9EgS6PqIBP2Z6zQB+PBRaJVT6/egK/X6E+rjXK+E/qI9xI8E3tUqQ9dHWKqHgjYSTVgnGrpGw0ysBmreDXgk2mMHSKsHS70c31O/HUn99PEFx0ithPT86lrmMag0/FvE9POAmS9Qw4dhljXs8UC4X4KYc//wGvrh5iy+rCGhK9MLmzXtmgCcn21toR9MO2EI/OqfH81MTtyOAN1niRzNrBYQxekcB53U2/7L4JwVrdL9pcx8x+0XurYBklUEoAj/OXLY4pRTgoxZ2HRYsW+bdyzlKBce1FlXzly1F+y32XlPFdn+RUQaAPZs3TJihVNC3pbKP/sW/nH9nx84CCKhEgB/ZDAUyB4e+UpIHIW/ws8+jN20Dl6PiTJ453PoyB30Kopq3iI+P3GOPgKpbV6W1jwImYVb3TmAi/HkSv73bO3v8nAeK54C/Jy5bHFkS696cPd+V8/Dt7K3qLqBIQzK9Psp6oKgiccA7XktW3Engov++cvk36RuHSZzdov5gKron+a+3S/0ucOu/FuojmzJf50onApgsuLdpx41sddWF1HrxwttcYHbahKpBm+ZHS3ZIA0MH776I90xdn4FA9QKT/OjKgjDUZj6dL67vAZONJaFdtowNE1DHWTbo2tiXjgyun4U4fwSUl8v6qPD58y7wO2Kkz612oRD89LuJqMrw0gTUL1CEo3705H/CxtuKIvL6uxgr8gYzeB35KxuOPJu2HMGh5Ci+2BZphE03QuDHesCPslcVm25jzGE/FmyCX5E4w36EPp/dN1XQnoBKP8q26qqKrgXMZoEiVsjZ/Zrst0V25Yq6s2ffQ+G3FRYYk1H4MfmiBEXX4qbl7jvst/3uCYIgCIIgCIIgCOIv8Q+cwCnnTSVOUwAAAABJRU5ErkJggg=="; // Replace with your Base64 encoded data
const imname = "testttttttttttttttt.jpg"; // Replace with the desired name
const type = "png"; // Replace with the file type

const mutation = `
  mutation MyMutation($base64str: String = "", $name: String = "", $type: String = "") {
    file_upload(base64str: $base64str, name: $name, type: $type) {
      image_url
    }
  }
`;

const variables = {
  base64str: base64str,
  name: imname,
  type: type,
};

fetch(graphqlUrl, {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
    "Accept": "application/json",
  },
  body: JSON.stringify({
    query: mutation,
    variables: variables,
  }),
})
  .then((response) => response.json())
  .then((data) => {
    console.log("Response:", data);
    const imageUrl = data.data.file_upload.image_url;
    console.log("Image URL:", imageUrl);
  })
  .catch((error) => {
    console.error("Error:", error);
  });
