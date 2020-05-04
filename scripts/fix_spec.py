import yaml

input = "./swagger/specification.yaml" # Which file to read
output = "./swagger/specification.yaml" # which file to write to


def Fix(obj: object, isDate: bool) -> object:
    if type(obj) != dict:
        return obj
    for k, v in obj.items():
        valIsDate = type(k) == str and k.lower().endswith("date")
        obj[k] = Fix(v, valIsDate)
        if isDate and k == "format":
            obj[k] = "date"
    return obj


if __name__ == '__main__':
    print("Started fixer...")
    parsed = None
    with open(input, "r") as file:
        parsed = yaml.safe_load(file)
    fixed = Fix(parsed, False)

    with open(output, "w") as file:
        yaml.dump(fixed, file)
    print("Fixer finished!")