words=123&uploadfile1=uploadfile1

POST http://httpbin.org/post HTTP/1.1
Content-Type: multipart/form-data; boundary=285fa365bd76e6378f91f09f4eae20877246bbba4d31370d3c87b752d350

multipart/form-data; boundary=285fa365bd76e6378f91f09f4eae20877246bbba4d31370d3c87b752d350
--285fa365bd76e6378f91f09f4eae20877246bbba4d31370d3c87b752d350
Content-Disposition: form-data; name="uploadFile1"; filename="uploadfile1.txt"
Content-Type: application/octet-stream

upload file1
--285fa365bd76e6378f91f09f4eae20877246bbba4d31370d3c87b752d350
Content-Disposition: form-data; name="uploadFile1"; filename="uploadfile2.txt"
Content-Type: application/octet-stream

upload file2
--285fa365bd76e6378f91f09f4eae20877246bbba4d31370d3c87b752d350
Content-Disposition: form-data; name="words"

123
--285fa365bd76e6378f91f09f4eae20877246bbba4d31370d3c87b752d350--
