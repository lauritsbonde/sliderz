<script lang="ts">
  let file: File | null = null;
  let startPage = 1;
  let endPage = 0;

  const handleFileChange = (event: Event) => {
    const target = event.target as HTMLInputElement;
    
    if(!target.files) {
      console.error("No file selected");
      return;
    }

    if(target.files.length > 1) {
      console.error("Please select only one file");
      return;
    }

    if(target.files.length === 0) {
      console.error("No file selected");
      return;
    }

    file = target.files[0];
    console.log(file);
  };

  const uploadFile = () => {
    if(!file) {
      console.error("No file selected");
      return;
    }

    const formData = new FormData();
    formData.append("file", file);
    formData.append("startPage", startPage?.toString() || "1");
    formData.append("endPage", endPage?.toString() || "1");

    fetch("http://localhost:8090/upload", {
      method: "POST",
      body: formData
    })
    .then(response => response.json())
    .then(data => {
      console.log(data);
    })
    .catch(error => {
      console.error(error);
    });
  }

</script>

<section class="container mx-auto">
  <h2 class="text-2xl font-bold">Upload your file</h2>
  <p class="text-lg">Select a file to convert it into presentation slides.</p>
  <input type="file" class="file-input file-input-bordered w-full max-w-xs" accept=".docx,.pdf,.txt" multiple={false}  on:change={handleFileChange}/>
  <p class="text-sm text-gray-100/25">Supported file types: .docx, .pdf, .txt</p>
  <div class="divider"></div>
  <div class="">
    <label for="startPage" class="text-sm">Start Page</label>
    <input type="number" id="startPage" class="input input-bordered w-full max-w-xs" bind:value={startPage} />

    <label for="endPage" class="text-sm">End Page</label>
    <input type="number" id="endPage" class="input input-bordered w-full max-w-xs" bind:value={endPage} />
  </div>
  <button class="btn btn-primary" on:click={uploadFile}>Upload!</button>
</section>