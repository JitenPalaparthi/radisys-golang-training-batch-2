#### 1- Standard packages        --> GOROOT/src
#### 2- User defined packages    --> Packages are created by the user/developer.
####								   2.0  dependency is the main issue.
####								   2.1. You can copy those to GOROOT/src
####								   2.2. You can add new path to GOPATH env --> bin/ pkg/ src/ [Old way of doing]
#### 								   2.3. With point 2.2, cannot maintain dependency

#### 3- Third party packages     --> Packages created by the community or other persons. same like point 2.

#### To resolve 2 and 3 go has a solution that is go mod
