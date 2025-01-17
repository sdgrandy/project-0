package main

import "os/exec"


func main(){
        // build image and container from Dockerfile

        // stop container
        dockerstop := exec.Command("sudo", "docker", "stop", "pgcontainer")
        dockerstop.Run()

        // remove container
        dockerdelete := exec.Command("sudo", "docker", "container", "rm", "pgcontainer")
        dockerdelete.Run()

        // remove image
        dockerrm := exec.Command("sudo","docker", "rmi", "pgimage")
        dockerrm.Run()

        // build image
        dockerbuild := exec.Command("sudo", "docker", "build", "-t", "pgimage", ".")
        dockerbuild.Run()

        // build container
        dockerrun := exec.Command("sudo", "docker", "run", "--name=pgcontainer", "-p=5432:5432", "-e POSTGRES_PASSWORD=password", "-d", "pgimage")
        dockerrun.Run()

        // // stop container
        // dockerstop = exec.Command("sudo", "docker", "stop", "pgcontainer")
        // dockerstop.Run()

        // // start container
        // dockerstart := exec.Command("sudo", "docker", "start", "pgcontainer")
        // dockerstart.Run()


        // to open psql shell- sudo docker exec -it pgcontainer psql -U postgres
        // to list tables in database in psql: \dt
        // to view rows in doctors table: select * from doctors;
        // to exit psql: \q
}