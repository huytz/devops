# devops

This is my solutions for `Test.md`, please check each directory with them same name of questions for more detail.

1. To allocate subnet for Nat Instance, I use interface call `allocator` and implement 2 type of allocate:
    - `allocate()` is normal allocated subnet for Nat instance.
    - `randomAllocate()` is a re-check after allocate has been run, if a instance has no subnet, then we randomly allocate.
    
    ```
    // Allocator allocate subnet
    type Allocator interface {
	    allocate()
	    randomAllocate()
    }
    ```

2. This is my workflow to get all Vulnerabilities from quay.io of image:

    - `Read json -> Image info -> Get Image manifest hash -> get all vulnerabilities.` 

3. For github team manage, I'm use https://registry.terraform.io/providers/integrations/github/latest/docs/resources/team_membership .