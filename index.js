// @ts-check

const _ = require("lodash");
const path = require("path");
const { name, version } = require("./package.json");

const httpUpdateMethodChoices = ["put", "patch"];

/** @type {import('caz').Template} */
module.exports = {
  name,
  version,
  prompts: [
    {
      name: "name",
      type: "text",
      message: "API resource name, e.g. article",
    },
    {
      name: "namespace",
      type: "text",
      message:
        "API namespace (leave empty to use resource name as namespace), e.g. blog",
    },
    {
      name: "api_version",
      type: "number",
      message: "API version, e.g. 1",
      initial: 1,
    },
    {
      name: "http_update_method",
      type: "select",
      message: "HTTP method for update action",
      choices: httpUpdateMethodChoices.map((v) => v.toUpperCase()),
    },
    {
      name: "go_module_name",
      type: "text",
      message: "Go module name for target app",
      initial: "app",
    },
  ],
  init: false,
  setup: async (ctx) => {
    // TODO: read from go.mod
    //ctx.answers.go_module_name = "app";
    ctx.answers.http_update_method =
      httpUpdateMethodChoices[ctx.answers.http_update_method] || "PUT";

    if (ctx.answers.namespace === "") {
      ctx.answers.namespace = ctx.answers.name;
    }

    const namespace = ctx.answers.namespace + "";
    if (namespace !== "") {
      ctx.answers.namespace_single = _.snakeCase(namespace)
        .toLowerCase()
        .replace(/_/g, "");
      ctx.answers.namespace_kebab = _.kebabCase(namespace);
      ctx.answers.namespace_upper = _.snakeCase(namespace).toUpperCase();
      ctx.answers.namespace_lower = _.snakeCase(namespace).toLowerCase();
      ctx.answers.namespace_pascal = _.upperFirst(_.camelCase(namespace));
    }

    const name = ctx.answers.name + "";
    if (name !== "") {
      ctx.answers.name_single = _.snakeCase(name)
        .toLowerCase()
        .replace(/_/g, "");
      ctx.answers.name_kebab = _.kebabCase(name);
      ctx.answers.name_upper = _.snakeCase(name).toUpperCase();
      ctx.answers.name_lower = _.snakeCase(name).toLowerCase();
      ctx.answers.name_pascal = _.upperFirst(_.camelCase(name));
    }
  },
  complete: async (ctx) => {
    // TODO: custom complete callback
    console.clear();
    console.log(
      `Created a new project in ${ctx.project} by the ${ctx.template} template.\n`,
    );
    console.log("Getting Started:");
    if (ctx.dest !== process.cwd()) {
      console.log(`  $ cd ${path.relative(process.cwd(), ctx.dest)}`);
    }
    console.log(`  $ make init`);
    console.log(`  $ make all`);
    console.log(`  $ make dev`);
    console.log("\nHappy hacking :)\n");
  },
};
